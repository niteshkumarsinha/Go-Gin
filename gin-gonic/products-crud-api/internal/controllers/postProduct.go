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
			resp := internal.NewHttpResponse(http.StatusBadRequest, err.Error())
			c.JSON(http.StatusBadRequest, resp)
			return
		}

		var guid = uuid.New().String()
		product.CreatedAt = time.Now().Format(time.RFC3339)

		if _, e := db.ExecContext(ctx, "INSERT INTO products (guid, name, price, description, created_at) VALUES (?, ?, ?, ?, ?)", guid, product.Name, product.Price, product.Description, product.CreatedAt); e != nil {
			resp := internal.NewHttpResponse(http.StatusInternalServerError, e.Error())
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		var productResponse internal.ProductResponse
		rows := db.QueryRow("SELECT guid, name, price, description, created_at FROM products WHERE guid = ?", guid)
		if e := rows.Scan(&productResponse.GUID, &productResponse.Name, &productResponse.Price, &productResponse.Description, &productResponse.CreatedAt); e != nil {
			resp := internal.NewHttpResponse(http.StatusInternalServerError, e.Error())
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		c.Writer.Header().Add("Location", "/products/"+guid)
		c.JSON(http.StatusCreated, internal.NewHttpResponse(http.StatusCreated, productResponse))
	}
}
