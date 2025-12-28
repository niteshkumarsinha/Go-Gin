package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitesh111sinha/products-crud-api/internal"
)

// GetProduct fetches a single product by its GUID
func GetProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		guid := c.Param("guid")
		var p internal.ProductResponse

		err := db.QueryRowContext(c.Request.Context(), "SELECT guid, name, price, description, created_at FROM products WHERE guid = ?", guid).
			Scan(&p.GUID, &p.Name, &p.Price, &p.Description, &p.CreatedAt)

		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, internal.NewHttpResponse(http.StatusNotFound, "Product not found"))
				return
			}
			c.JSON(http.StatusInternalServerError, internal.NewHttpResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, internal.NewHttpResponse(http.StatusOK, p))
	}
}
