package main

import (
	"goplace/configs"
	"goplace/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//Run database
	configs.ConnectionDatabase()

	//Instance HTTP server
	http := gin.Default()

	//Routes
	routes.ProductRoute(http)
	routes.PriceRoute(http)

	//Run
	http.Run("0.0.0.0:6000")
}
