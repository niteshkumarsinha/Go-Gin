package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nitesh111sinha/products-crud-api/internal"
)

func PostProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product internal.Product
		ctx := c.Request.Context()

		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, internal.NewHttpResponse(http.StatusBadRequest, err.Error()))
			return
		}

		var guid = uuid.New().String()
		product.CreatedAt = time.Now().Format(time.RFC3339)

		if _, e := db.ExecContext(ctx, "INSERT INTO products (guid, name, price, description, created_at) VALUES (?, ?, ?, ?, ?)", guid, product.Name, product.Price, product.Description, product.CreatedAt); e != nil {
			c.JSON(http.StatusInternalServerError, internal.NewHttpResponse(http.StatusInternalServerError, e.Error()))
			return
		}

		var productResponse internal.ProductResponse
		if err := db.QueryRowContext(ctx, "SELECT guid, name, price, description, created_at FROM products WHERE guid = ?", guid).Scan(&productResponse.GUID, &productResponse.Name, &productResponse.Price, &productResponse.Description, &productResponse.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, internal.NewHttpResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.Writer.Header().Add("Location", "/products/"+guid)
		c.JSON(http.StatusCreated, internal.NewHttpResponse(http.StatusCreated, productResponse))
	}
}
