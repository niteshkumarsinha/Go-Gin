package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get a product by ID
// @Tags Products
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /products/{id} [get]
func GetProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product fetched successfully",
		})
	}
}
