package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"ip_area_query/config"
	"log"
	"time"
)

var (
	db *gorm.DB
)

func InitMysql() {

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := config.C.Mysql.Dsn
	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		log.Fatal(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.C.Mysql.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.C.Mysql.OpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(config.C.Mysql.ConnMaxLifetime))
	db = gdb
	db.Logger.LogMode(logger.Error)
}
