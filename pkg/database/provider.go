package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func ProviderGorm() *gorm.DB {
	//db, err := gorm.Open("mysql", "root:123456@(192.168.110.129:3337)/video?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := "root:123456@(192.168.110.130:3337)/video?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 迁移 schema
	//db.AutoMigrate(&database.VideoInfo{}, &database.PlayInfo{})
	//db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(&database.VideoInfo{}, &database.PlayInfo{})
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
