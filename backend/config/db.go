package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func DatabaseInit() {
	host := "localhost"
	user := "root"
	password := "password"
	dbName := "app"
	port := "3306"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return database
}
