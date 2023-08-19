package routes

import (
	"goplace/controllers"

	"github.com/gin-gonic/gin"
)

func PriceRoute(router *gin.Engine) {
	router.GET("/price", controllers.FindAllPrice)
	router.GET("/price/:product_id", controllers.FindPriceByProductId)
	router.POST("/price", controllers.AddPrice)
}
