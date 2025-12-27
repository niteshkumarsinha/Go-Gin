package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// @Summary Delete a product by ID
// @Tags Products
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}