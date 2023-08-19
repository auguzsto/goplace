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

	if err := configs.DB.Preload("Product").Find(&prices).Error; err != nil {
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

	if err := configs.DB.Preload("Product").Where("product_id = ?", product.ID).First(&price).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, price)
}

func AddPrice(c *gin.Context) {
	var product models.Product
	var price models.Price
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
		QuantityMost: dto.QuantityMost,
		Price:        dto.Price,
		PriceMost:    dto.PriceMost,
		PriceOffer:   dto.PriceOffer,
	}).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	configs.DB.Preload("Product").Find(&price)

	c.IndentedJSON(http.StatusOK, price)

}

func UpdatePriceByProductId(c *gin.Context) {
	var price models.Price
	var dto dto.PriceDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := configs.DB.Model(&price).Where("product_id = ?", c.Param("product_id")).Updates(&models.Price{
		QuantityMost: dto.QuantityMost,
		Price:        dto.Price,
		PriceMost:    dto.PriceMost,
		PriceOffer:   dto.PriceOffer,
	}).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	configs.DB.Preload("Product").Where("product_id = ?", c.Param("product_id")).First(&price)

	c.IndentedJSON(http.StatusOK, price)
}
