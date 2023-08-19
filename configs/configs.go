package configs

import (
	"goplace/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	dns := "root:password@tcp(127.0.0.1:3306)/goplace?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//Run migrations
	database.AutoMigrate(&models.Product{})

	DB = database
}
