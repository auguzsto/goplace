package controllers

import (
	"goplace/configs"
	"goplace/dto"
	"goplace/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FindAllProducts(c *gin.Context) {
	var products []models.Product

	configs.DB.Find(&products)

	c.IndentedJSON(http.StatusOK, products)
}

func FindProductById(c *gin.Context) {
	var product models.Product

	if err := configs.DB.First(&product, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}

func AddProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		return
	}

	configs.DB.Create(&models.Product{
		Code:     product.Code,
		Quantity: product.Quantity,
	})

	configs.DB.Find(&product)

	c.IndentedJSON(http.StatusOK, product)
}

func UpdateProductById(c *gin.Context) {

	var product models.Product
	var dto dto.ProductDTO

	if err := configs.DB.First(&product, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := configs.DB.Model(&product).Updates(&models.Product{Code: dto.Code, Quantity: dto.Quantity}).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, product)

}

func DeleteProductById(c *gin.Context) {
	var product models.Product
	time := time.Now()

	if err := configs.DB.First(&product, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := configs.DB.Model(&product).Updates(&models.Product{
		DeletedAt: time,
	}).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	configs.DB.First(&product)

	c.IndentedJSON(http.StatusOK, product)

}
