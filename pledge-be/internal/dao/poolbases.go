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

var _ PoolbasesDao = (*poolbasesDao)(nil)

// PoolbasesDao 资金池数据访问接口
type PoolbasesDao interface {
	// Create 创建资金池记录
	Create(ctx context.Context, table *model.Poolbases) error
	// DeleteByID 根据 ID 删除资金池
	DeleteByID(ctx context.Context, id uint64) error
	// UpdateByID 根据 ID 更新资金池（支持部分更新）
	UpdateByID(ctx context.Context, table *model.Poolbases) error
	// GetByID 根据 ID 获取资金池
	GetByID(ctx context.Context, id uint64) (*model.Poolbases, error)
	// GetByColumns 按条件分页查询资金池列表
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.Poolbases, int64, error)

	// CreateByTx 在事务中创建资金池记录
	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.Poolbases) (uint64, error)
	// DeleteByTx 在事务中根据 ID 删除资金池
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	// UpdateByTx 在事务中根据 ID 更新资金池
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Poolbases) error
}

type poolbasesDao struct {
	db    *gorm.DB
	cache cache.PoolbasesCache // if nil, the cache is not used.
	sfg   *singleflight.Group  // if cache is nil, the sfg is not used.
}

// NewPoolbasesDao 创建资金池数据访问实例，可选传入缓存以实现缓存读写
func NewPoolbasesDao(db *gorm.DB, xCache cache.PoolbasesCache) PoolbasesDao {
	if xCache == nil {
		return &poolbasesDao{db: db}
	}
	return &poolbasesDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

// deleteCache 删除资金池缓存（如果缓存实例不为 nil）
func (d *poolbasesDao) deleteCache(ctx context.Context, id uint64) error {
	if d.cache != nil {
		return d.cache.Del(ctx, id)
	}
	return nil
}

// Create 创建资金池记录，插入后 ID 值会回写到传入的 table 参数中
func (d *poolbasesDao) Create(ctx context.Context, table *model.Poolbases) error {
	return d.db.WithContext(ctx).Create(table).Error
}

// DeleteByID 根据 ID 删除资金池记录，同时清除对应的缓存
func (d *poolbasesDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Poolbases{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByID 根据 ID 更新资金池记录（只更新非零字段），同时清除缓存
func (d *poolbasesDao) UpdateByID(ctx context.Context, table *model.Poolbases) error {
	err := d.updateDataByID(ctx, d.db, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}

// updateDataByID 根据 ID 更新资金池数据，只将非空字段加入更新 map
func (d *poolbasesDao) updateDataByID(ctx context.Context, db *gorm.DB, table *model.Poolbases) error {
	if table.ID < 1 {
		return errors.New("id cannot be 0")
	}

	update := map[string]interface{}{}

	if table.PoolID != 0 {
		update["pool_id"] = table.PoolID
	}
	if table.SettleTime != "" {
		update["settle_time"] = table.SettleTime
	}
	if table.EndTime != "" {
		update["end_time"] = table.EndTime
	}
	if table.InterestRate != "" {
		update["interest_rate"] = table.InterestRate
	}
	if table.MaxSupply != "" {
		update["max_supply"] = table.MaxSupply
	}
	if table.LendSupply != "" {
		update["lend_supply"] = table.LendSupply
	}
	if table.BorrowSupply != "" {
		update["borrow_supply"] = table.BorrowSupply
	}
	if table.MortgageRate != "" {
		update["mortgage_rate"] = table.MortgageRate
	}
	if table.LendToken != "" {
		update["lend_token"] = table.LendToken
	}
	if table.BorrowToken != "" {
		update["borrow_token"] = table.BorrowToken
	}
	if table.State != "" {
		update["state"] = table.State
	}
	if table.LendDebtToken != "" {
		update["lend_debt_token"] = table.LendDebtToken
	}
	if table.BorrowDebtToken != "" {
		update["borrow_debt_token"] = table.BorrowDebtToken
	}
	if table.AutoLiquidateThreshold != "" {
		update["auto_liquidate_threshold"] = table.AutoLiquidateThreshold
	}
	if table.BorrowTokenInfo.String() != "" {
		update["borrow_token_info"] = table.BorrowTokenInfo
	}
	if table.LendTokenInfo.String() != "" {
		update["lend_token_info"] = table.LendTokenInfo
	}
	if table.ChainID != "" {
		update["chain_id"] = table.ChainID
	}
	if table.LendTokenSymbol != "" {
		update["lend_token_symbol"] = table.LendTokenSymbol
	}
	if table.BorrowTokenSymbol != "" {
		update["borrow_token_symbol"] = table.BorrowTokenSymbol
	}

	return db.WithContext(ctx).Model(table).Updates(update).Error
}

// GetByID 根据 ID 获取资金池，优先从缓存读取，缓存未命中则查数据库并回写缓存，使用 singleflight 防止缓存击穿
func (d *poolbasesDao) GetByID(ctx context.Context, id uint64) (*model.Poolbases, error) {
	// no cache
	if d.cache == nil {
		record := &model.Poolbases{}
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
			table := &model.Poolbases{}
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
			if err = d.cache.Set(ctx, id, table, cache.PoolbasesExpireTime); err != nil {
				logger.Warn("cache.Set error", logger.Err(err), logger.Any("id", id))
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.Poolbases)
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

// GetByColumns 按自定义条件分页查询资金池列表，支持排序和分页参数。
// 详见 https://go-sponge.com/component/data/custom-page-query.html
func (d *poolbasesDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.Poolbases, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions(query.WithWhitelistNames(model.PoolbasesColumnNames))
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.Poolbases{}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.Poolbases{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// CreateByTx 在给定事务中创建资金池记录，返回记录 ID
func (d *poolbasesDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.Poolbases) (uint64, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.ID, err
}

// DeleteByTx 在给定事务中根据 ID 删除资金池记录，同时清除缓存
func (d *poolbasesDao) DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Where("id = ?", id).Delete(&model.Poolbases{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByTx 在给定事务中根据 ID 更新资金池记录，同时清除缓存
func (d *poolbasesDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Poolbases) error {
	err := d.updateDataByID(ctx, tx, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}
