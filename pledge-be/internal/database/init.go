// Package database provides database client initialization.
package database

import (
	"strings"
	"sync"

	"github.com/go-dev-frame/sponge/pkg/sgorm"

	"pledge-be/internal/config"
)

var (
	gdb     *sgorm.DB
	gdbOnce sync.Once

	// ErrRecordNotFound 记录未找到错误
	ErrRecordNotFound = sgorm.ErrRecordNotFound
)

// InitDB 初始化数据库连接，根据配置选择 MySQL 或 TiDB 驱动
func InitDB() {
	dbDriver := config.Get().Database.Driver
	switch strings.ToLower(dbDriver) {
	case sgorm.DBDriverMysql, sgorm.DBDriverTidb:
		gdb = InitMysql()
	default:
		panic("InitDB error, please modify the correct 'database' configuration at yaml file. " +
			"Refer to https://pledge-be/blob/main/configs/pledge_be.yml#L85")
	}
}

// GetDB 获取数据库连接实例（单例，懒加载）
func GetDB() *sgorm.DB {
	if gdb == nil {
		gdbOnce.Do(func() {
			InitDB()
		})
	}

	return gdb
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	return sgorm.CloseDB(gdb)
}
