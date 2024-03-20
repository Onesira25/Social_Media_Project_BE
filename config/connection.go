package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitSQL(c AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("terjadi error", err.Error())
		return nil
	}

	return db
}

func Migrate(connection *gorm.DB, dst ...interface{}) error {
	return connection.AutoMigrate(dst...)
}
