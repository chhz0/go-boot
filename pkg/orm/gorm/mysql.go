package gorm

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsnFormat = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

type MysqlOptions struct {
	Host     string
	User     string
	Password string
	Database string

	MaxIdleConns    int
	MaxOpenConns    int
	MaxConnLifetime time.Duration
	MaxIdleTime     time.Duration
	// EnableDebug enables debug mode for GORM.
	EnableDebug bool

	AutoMigrate       bool
	AutoMigrateTables []any
}

func (o *MysqlOptions) MysqlConnect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(o.Dsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if o.AutoMigrate {
		err := db.AutoMigrate(o.AutoMigrateTables...)
		if err != nil {
			return nil, err
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(o.MaxIdleConns)
	sqlDB.SetMaxOpenConns(o.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(o.MaxConnLifetime)
	sqlDB.SetConnMaxIdleTime(o.MaxIdleTime)

	return db, nil
}

func (o *MysqlOptions) Dsn() string {
	return fmt.Sprintf(dsnFormat, o.User, o.Password, o.Host, o.Database)
}

func NewMysqlOptions(host, user, password, database string) *MysqlOptions {
	return &MysqlOptions{
		Host:     host,
		User:     user,
		Password: password,
		Database: database,

		MaxIdleConns:    10,
		MaxOpenConns:    100,
		MaxConnLifetime: time.Second * 30,
		MaxIdleTime:     time.Second * 30,

		EnableDebug: false,

		AutoMigrate:       false,
		AutoMigrateTables: make([]any, 0),
	}
}
