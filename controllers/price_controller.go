package controllers

import (
	"goplace/configs"
	"goplace/dto"
	"goplace/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllPrice(c *gin.Context) {
	var prices []models.Price

	if err := configs.DB.Find(&prices).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, prices)
}

func FindPriceByProductId(c *gin.Context) {
	var price models.Price
	var product models.Product

	if err := configs.DB.First(&product, c.Param("product_id")).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := configs.DB.Where("product_id = ?", product.ID).First(&price).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, price)
}

func AddPrice(c *gin.Context) {
	var product models.Product
	var dto dto.PriceDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := configs.DB.First(&product, dto.ProductID).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := configs.DB.Create(&models.Price{
		Product:      product,
		ProductID:    product.ID,
		ProductCode:  product.Code,
		QuantityMost: dto.QuantityMost,
		Price:        dto.Price,
		PriceMost:    dto.PriceMost,
		PriceOffer:   dto.PriceOffer,
	}).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, dto)

}
