package main

import (
	"goplace/configs"
	"goplace/models"
	"goplace/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//Run database with migrations
	db := configs.DbCon()
	db.AutoMigrate(&models.Product{}, &models.Stock{})

	//Instance HTTP server
	http := gin.Default()

	//Routes
	routes.ProductRoute(http, db)

	//Run
	http.Run("0.0.0.0:6000")
}
