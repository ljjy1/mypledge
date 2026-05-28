package dao

import (
	"context"
	"errors"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"pledge-be/internal/cache"
	"pledge-be/internal/database"
	"pledge-be/internal/model"
)

var _ UserDao = (*userDao)(nil)

// UserDao 用户数据访问接口
type UserDao interface {
	// Create 创建用户记录
	Create(ctx context.Context, table *model.User) error
	// DeleteByID 根据 ID 删除用户
	DeleteByID(ctx context.Context, id uint64) error
	// UpdateByID 根据 ID 更新用户（支持部分更新）
	UpdateByID(ctx context.Context, table *model.User) error
	// GetByID 根据 ID 获取用户
	GetByID(ctx context.Context, id uint64) (*model.User, error)
	// GetByColumns 按条件分页查询用户列表
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.User, int64, error)
	// GetByLogin 根据登录名获取用户
	GetByLogin(ctx context.Context, login string) (*model.User, error)

	// CreateByTx 在事务中创建用户记录
	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.User) (uint64, error)
	// DeleteByTx 在事务中根据 ID 删除用户
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	// UpdateByTx 在事务中根据 ID 更新用户
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.User) error
}

type userDao struct {
	db    *gorm.DB
	cache cache.UserCache     // if nil, the cache is not used.
	sfg   *singleflight.Group // if cache is nil, the sfg is not used.
}

// NewUserDao 创建用户数据访问实例，可选传入缓存以实现缓存读写
func NewUserDao(db *gorm.DB, xCache cache.UserCache) UserDao {
	if xCache == nil {
		return &userDao{db: db}
	}
	return &userDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

// deleteCache 删除用户缓存（如果缓存实例不为 nil）
func (d *userDao) deleteCache(ctx context.Context, id uint64) error {
	if d.cache != nil {
		return d.cache.Del(ctx, id)
	}
	return nil
}

// Create 创建用户记录，插入后 ID 值会回写到传入的 table 参数中
func (d *userDao) Create(ctx context.Context, table *model.User) error {
	return d.db.WithContext(ctx).Create(table).Error
}

// DeleteByID 根据 ID 删除用户记录，同时清除对应的缓存
func (d *userDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByID 根据 ID 更新用户记录（只更新非零字段），同时清除缓存
func (d *userDao) UpdateByID(ctx context.Context, table *model.User) error {
	err := d.updateDataByID(ctx, d.db, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}

// updateDataByID 根据 ID 更新用户数据，只将非空字段加入更新 map
func (d *userDao) updateDataByID(ctx context.Context, db *gorm.DB, table *model.User) error {
	if table.ID < 1 {
		return errors.New("id cannot be 0")
	}

	update := map[string]interface{}{}

	if table.Login != "" {
		update["login"] = table.Login
	}
	if table.Nike != "" {
		update["nike"] = table.Nike
	}
	if table.Password != "" {
		update["password"] = table.Password
	}

	return db.WithContext(ctx).Model(table).Updates(update).Error
}

// GetByID 根据 ID 获取用户，优先从缓存读取，缓存未命中则查数据库并回写缓存，使用 singleflight 防止缓存击穿
func (d *userDao) GetByID(ctx context.Context, id uint64) (*model.User, error) {
	// no cache
	if d.cache == nil {
		record := &model.User{}
		err := d.db.WithContext(ctx).Where("id = ?", id).First(record).Error
		return record, err
	}

	// get from cache
	record, err := d.cache.Get(ctx, id)
	if err == nil {
		return record, nil
	}

	// get from database
	if errors.Is(err, database.ErrCacheNotFound) {
		// for the same id, prevent high concurrent simultaneous access to database
		val, err, _ := d.sfg.Do(utils.Uint64ToStr(id), func() (interface{}, error) { //nolint
			table := &model.User{}
			err = d.db.WithContext(ctx).Where("id = ?", id).First(table).Error
			if err != nil {
				if errors.Is(err, database.ErrRecordNotFound) {
					// set placeholder cache to prevent cache penetration, default expiration time 10 minutes
					if err = d.cache.SetPlaceholder(ctx, id); err != nil {
						logger.Warn("cache.SetPlaceholder error", logger.Err(err), logger.Any("id", id))
					}
					return nil, database.ErrRecordNotFound
				}
				return nil, err
			}
			// set cache
			if err = d.cache.Set(ctx, id, table, cache.UserExpireTime); err != nil {
				logger.Warn("cache.Set error", logger.Err(err), logger.Any("id", id))
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.User)
		if !ok {
			return nil, database.ErrRecordNotFound
		}
		return table, nil
	}

	if d.cache.IsPlaceholderErr(err) {
		return nil, database.ErrRecordNotFound
	}

	return nil, err
}

// GetByColumns 按自定义条件分页查询用户列表，支持排序和分页参数。
// 详见 https://go-sponge.com/component/data/custom-page-query.html
func (d *userDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.User, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions(query.WithWhitelistNames(model.UserColumnNames))
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.User{}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.User{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// GetByLogin 根据登录名获取用户
func (d *userDao) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	record := &model.User{}
	err := d.db.WithContext(ctx).Where("login = ?", login).First(record).Error
	return record, err
}

// CreateByTx 在给定事务中创建用户记录，返回记录 ID
func (d *userDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.User) (uint64, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.ID, err
}

// DeleteByTx 在给定事务中根据 ID 删除用户记录，同时清除缓存
func (d *userDao) DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByTx 在给定事务中根据 ID 更新用户记录，同时清除缓存
func (d *userDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.User) error {
	err := d.updateDataByID(ctx, tx, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}
