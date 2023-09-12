package database

import (
	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewConnection() (Connection, error) {
	gormParams := gormMysql.Open(buildDBConfig().FormatDSN())
	db, err := gorm.Open(gormParams, &gorm.Config{})
	return Connection{
		Read:  db,
		Write: db,
	}, err
}

func buildDBConfig() *mysql.Config {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "root"
	cfg.DBName = "gcommerce"
	return cfg
}
