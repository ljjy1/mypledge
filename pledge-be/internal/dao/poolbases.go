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
	Create(ctx context.Context, table *model.Poolbases) error
	DeleteByID(ctx context.Context, id uint64) error
	UpdateByID(ctx context.Context, table *model.Poolbases) error
	GetByID(ctx context.Context, id uint64) (*model.Poolbases, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.Poolbases, int64, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.Poolbases) (uint64, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Poolbases) error
}

type poolbasesDao struct {
	db    *gorm.DB
	cache cache.PoolbasesCache
	sfg   *singleflight.Group
}

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

func (d *poolbasesDao) deleteCache(ctx context.Context, id uint64) error {
	if d.cache != nil {
		return d.cache.Del(ctx, id)
	}
	return nil
}

func (d *poolbasesDao) Create(ctx context.Context, table *model.Poolbases) error {
	return d.db.WithContext(ctx).Create(table).Error
}

func (d *poolbasesDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Poolbases{}).Error
	if err != nil {
		return err
	}
	_ = d.deleteCache(ctx, id)
	return nil
}

func (d *poolbasesDao) UpdateByID(ctx context.Context, table *model.Poolbases) error {
	err := d.updateDataByID(ctx, d.db, table)
	_ = d.deleteCache(ctx, table.ID)
	return err
}

func (d *poolbasesDao) updateDataByID(ctx context.Context, db *gorm.DB, table *model.Poolbases) error {
	if table.ID < 1 {
		return errors.New("id cannot be 0")
	}

	update := map[string]interface{}{}

	if table.ContractID != 0 {
		update["contract_id"] = table.ContractID
	}
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
	if table.ChainID != "" {
		update["chain_id"] = table.ChainID
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

func (d *poolbasesDao) GetByID(ctx context.Context, id uint64) (*model.Poolbases, error) {
	if d.cache == nil {
		record := &model.Poolbases{}
		err := d.db.WithContext(ctx).Where("id = ?", id).First(record).Error
		return record, err
	}

	record, err := d.cache.Get(ctx, id)
	if err == nil {
		return record, nil
	}

	if errors.Is(err, database.ErrCacheNotFound) {
		val, err, _ := d.sfg.Do(utils.Uint64ToStr(id), func() (interface{}, error) {
			table := &model.Poolbases{}
			err = d.db.WithContext(ctx).Where("id = ?", id).First(table).Error
			if err != nil {
				if errors.Is(err, database.ErrRecordNotFound) {
					if err = d.cache.SetPlaceholder(ctx, id); err != nil {
						logger.Warn("cache.SetPlaceholder error", logger.Err(err), logger.Any("id", id))
					}
					return nil, database.ErrRecordNotFound
				}
				return nil, err
			}
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

func (d *poolbasesDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.Poolbases, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions(query.WithWhitelistNames(model.PoolbasesColumnNames))
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" {
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

func (d *poolbasesDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.Poolbases) (uint64, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.ID, err
}

func (d *poolbasesDao) DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Where("id = ?", id).Delete(&model.Poolbases{}).Error
	if err != nil {
		return err
	}
	_ = d.deleteCache(ctx, id)
	return nil
}

func (d *poolbasesDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.Poolbases) error {
	err := d.updateDataByID(ctx, tx, table)
	_ = d.deleteCache(ctx, table.ID)
	return err
}
