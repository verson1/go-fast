package db

import (
	"errors"
	"go-fast/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

const (
	MysqlDriver   = "mysql"
	PostgreDriver = "postgre"
	SqliteDiver   = "sqlite"
)

func NewDB(config *config.DBConfig) (db *gorm.DB, err error) {
	if config == nil {
		panic("not init db config")
	}
	switch config.Driver {
	case MysqlDriver:
		db, err = gorm.Open(mysql.Open(config.Dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.LogLevel(config.LogLever))})
	case PostgreDriver:
		db, err = gorm.Open(postgres.Open(config.Dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.LogLevel(config.LogLever))})
	case SqliteDiver:
		db, err = gorm.Open(sqlite.Open(config.Dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.LogLevel(config.LogLever))})
	default:
		return nil, errors.New("db: unsupported db driver: " + config.Driver)
	}
	sqlDb, err := db.DB()
	if err != nil {
		return
	}
	// 设置连接池参数
	sqlDb.SetMaxIdleConns(config.MaxIdleConn)
	sqlDb.SetMaxOpenConns(config.MaxOpenConn)
	sqlDb.SetConnMaxIdleTime(time.Second * time.Duration(config.ConnMaxLifetime))
	return
}
