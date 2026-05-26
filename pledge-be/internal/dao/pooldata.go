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

var _ PooldataDao = (*pooldataDao)(nil)

// PooldataDao defining the dao interface
type PooldataDao interface {
	Create(ctx context.Context, table *model.Pooldata) error
	DeleteByID(ctx context.Context, id uint64) error
	UpdateByID(ctx context.Context, table *model.Pooldata) error
	GetByID(ctx context.Context, id uint64) (*model.Pooldata, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.Pooldata, int64, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.Pooldata) (uint64, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Pooldata) error
}

type pooldataDao struct {
	db    *gorm.DB
	cache cache.PooldataCache // if nil, the cache is not used.
	sfg   *singleflight.Group // if cache is nil, the sfg is not used.
}

// NewPooldataDao creating the dao interface
func NewPooldataDao(db *gorm.DB, xCache cache.PooldataCache) PooldataDao {
	if xCache == nil {
		return &pooldataDao{db: db}
	}
	return &pooldataDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

func (d *pooldataDao) deleteCache(ctx context.Context, id uint64) error {
	if d.cache != nil {
		return d.cache.Del(ctx, id)
	}
	return nil
}

// Create a new pooldata, insert the record and the id value is written back to the table
func (d *pooldataDao) Create(ctx context.Context, table *model.Pooldata) error {
	return d.db.WithContext(ctx).Create(table).Error
}

// DeleteByID delete a pooldata by id
func (d *pooldataDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Pooldata{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByID update a pooldata by id, support partial update
func (d *pooldataDao) UpdateByID(ctx context.Context, table *model.Pooldata) error {
	err := d.updateDataByID(ctx, d.db, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}

func (d *pooldataDao) updateDataByID(ctx context.Context, db *gorm.DB, table *model.Pooldata) error {
	if table.ID < 1 {
		return errors.New("id cannot be 0")
	}

	update := map[string]interface{}{}

	if table.ChainID != "" {
		update["chain_id"] = table.ChainID
	}
	if table.PoolID != "" {
		update["pool_id"] = table.PoolID
	}
	if table.SettleAmountLend != "" {
		update["settle_amount_lend"] = table.SettleAmountLend
	}
	if table.SettleAmountBorrow != "" {
		update["settle_amount_borrow"] = table.SettleAmountBorrow
	}
	if table.FinishAmountLend != "" {
		update["finish_amount_lend"] = table.FinishAmountLend
	}
	if table.FinishAmountBorrow != "" {
		update["finish_amount_borrow"] = table.FinishAmountBorrow
	}
	if table.LiquidationAmounLend != "" {
		update["liquidation_amoun_lend"] = table.LiquidationAmounLend
	}
	if table.LiquidationAmounBorrow != "" {
		update["liquidation_amoun_borrow"] = table.LiquidationAmounBorrow
	}

	return db.WithContext(ctx).Model(table).Updates(update).Error
}

// GetByID get a pooldata by id
func (d *pooldataDao) GetByID(ctx context.Context, id uint64) (*model.Pooldata, error) {
	// no cache
	if d.cache == nil {
		record := &model.Pooldata{}
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
			table := &model.Pooldata{}
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
			if err = d.cache.Set(ctx, id, table, cache.PooldataExpireTime); err != nil {
				logger.Warn("cache.Set error", logger.Err(err), logger.Any("id", id))
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.Pooldata)
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

// GetByColumns get a paginated list of pooldatas by custom conditions.
// For more details, please refer to https://go-sponge.com/component/data/custom-page-query.html
func (d *pooldataDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.Pooldata, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions(query.WithWhitelistNames(model.PooldataColumnNames))
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.Pooldata{}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.Pooldata{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// CreateByTx create a record in the database using the provided transaction
func (d *pooldataDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.Pooldata) (uint64, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.ID, err
}

// DeleteByTx delete a record by id in the database using the provided transaction
func (d *pooldataDao) DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Where("id = ?", id).Delete(&model.Pooldata{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByTx update a record by id in the database using the provided transaction
func (d *pooldataDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Pooldata) error {
	err := d.updateDataByID(ctx, tx, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}
