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
		rows, err := db.Query("SELECT guid, name, price, description, created_at FROM products")
		if err != nil {
			resp := internal.NewHttpResponse(http.StatusInternalServerError, err.Error())
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		defer rows.Close()

		products := []internal.ProductResponse{}
		for rows.Next() {
			var p internal.ProductResponse
			if err := rows.Scan(&p.GUID, &p.Name, &p.Price, &p.Description, &p.CreatedAt); err != nil {
				resp := internal.NewHttpResponse(http.StatusInternalServerError, err.Error())
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
			products = append(products, p)
		}

		if len(products) == 0 {
			resp := internal.NewHttpResponse(http.StatusNotFound, "No products found")
			c.JSON(http.StatusNotFound, resp)
			return
		}

		resp := internal.NewHttpResponse(http.StatusOK, products)
		c.JSON(http.StatusOK, resp)
	}
}
