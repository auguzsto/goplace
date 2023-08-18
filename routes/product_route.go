package routes

import (
	"goplace/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {
	router.GET("/products", controllers.FindAllProducts)
	router.GET("/products/:id", controllers.FindByIdProduct)
	router.POST("/products", controllers.AddProduct)
	router.PATCH("/products/:id", controllers.UpdateById)
	router.DELETE("/products/:id", controllers.DeleteById)
}
