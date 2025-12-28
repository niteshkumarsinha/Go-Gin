package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitesh111sinha/products-crud-api/internal"
)

// GetProducts fetches all products from the database
func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.QueryContext(c.Request.Context(), "SELECT guid, name, price, description, created_at FROM products")
		if err != nil {
			c.JSON(http.StatusInternalServerError, internal.NewHttpResponse(http.StatusInternalServerError, err.Error()))
			return
		}
		defer rows.Close()

		products := []internal.ProductResponse{}
		for rows.Next() {
			var p internal.ProductResponse
			if err := rows.Scan(&p.GUID, &p.Name, &p.Price, &p.Description, &p.CreatedAt); err != nil {
				c.JSON(http.StatusInternalServerError, internal.NewHttpResponse(http.StatusInternalServerError, err.Error()))
				return
			}
			products = append(products, p)
		}

		c.JSON(http.StatusOK, internal.NewHttpResponse(http.StatusOK, products))
	}
}
