package controllers

import (
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
)

// @Summary Update a product by ID
// @Tags Products
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /products/{id} [put]
func PutProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product updated successfully",
		})
	}
}
