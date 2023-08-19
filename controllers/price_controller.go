package controllers

import (
	"goplace/configs"
	"goplace/dto"
	"goplace/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
