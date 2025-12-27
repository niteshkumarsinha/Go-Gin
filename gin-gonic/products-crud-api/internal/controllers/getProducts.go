package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all products
// @Tags Products
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /products [get]
func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Products fetched successfully",
		})
	}
}