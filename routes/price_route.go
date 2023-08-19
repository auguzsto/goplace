package routes

import (
	"goplace/controllers"

	"github.com/gin-gonic/gin"
)

func PriceRoute(router *gin.Engine) {
	router.POST("/price", controllers.AddPrice)
}
