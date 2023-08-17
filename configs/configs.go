package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbCon() *gorm.DB {
	dns := "root:password@tcp(127.0.0.1:3306)/goplace?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
