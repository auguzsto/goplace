package controllers

import (
	"goplace/dto"
	"goplace/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		id := c.Query("id")

		if len(id) > 0 {
			db.First(&products, id)
			c.IndentedJSON(200, products)
			return
		}

		db.Find(&products)

		c.IndentedJSON(200, products)
	}
}

func AddProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product

		if err := c.ShouldBindJSON(&product); err != nil {
			return
		}

		db.Create(&models.Product{
			Code:  product.Code,
			Price: product.Price,
		})

		c.IndentedJSON(200, product)
	}

}

func UpdateById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		var dto dto.ProductDTO
		id := c.Query("id")

		if err := c.ShouldBindJSON(&dto); err != nil {
			return
		}

		db.First(&product, id)

		db.Model(&product).Updates(models.Product{
			Code:  dto.Code,
			Price: dto.Price,
		})

		c.IndentedJSON(200, product)
	}
}

func DeleteById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		id := c.Query("id")

		db.Delete(&product, id)

		c.IndentedJSON(200, product)
	}
}
