package routes

import (
	"goplace/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoute(router *gin.Engine, db *gorm.DB) {
	router.GET("/products", controllers.GetProduct(db))
	router.POST("/products", controllers.AddProduct(db))
	router.PATCH("/products", controllers.UpdateById(db))
	router.DELETE("/products", controllers.DeleteById(db))
}
