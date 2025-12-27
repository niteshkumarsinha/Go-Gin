package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new product
// @Tags Products
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /products [post]
func PostProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product created successfully",
		})
	}
}
