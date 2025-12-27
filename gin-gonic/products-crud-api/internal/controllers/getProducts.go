package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// @Summary Get all products
// @Tags Products
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /products [get]
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Products fetched successfully",
	})
}