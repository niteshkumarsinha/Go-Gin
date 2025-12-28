package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitesh111sinha/products-crud-api/internal"
)

type GUIDBinding struct {
	GUID string `uri:"guid" binding:"required,uuid4"`
}

// GetProduct fetches a single product by its GUID
func GetProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GUIDBinding

		if err := c.ShouldBindUri(&req); err != nil {
			resp := internal.NewHttpResponse(http.StatusBadRequest, err.Error())
			c.JSON(http.StatusBadRequest, resp)
			return
		}

		guid := req.GUID
		var p internal.ProductResponse
		ctx := c.Request.Context()
		rows := db.QueryRowContext(ctx, "SELECT guid, name, price, description, created_at FROM products WHERE guid=?", guid)
		if err := rows.Scan(&p.GUID, &p.Name, &p.Price, &p.Description, &p.CreatedAt); err != nil {
			if err == sql.ErrNoRows {
				resp := internal.NewHttpResponse(http.StatusNotFound, "Product not found")
				c.JSON(http.StatusNotFound, resp)
				return
			}
			resp := internal.NewHttpResponse(http.StatusInternalServerError, err.Error())
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		c.JSON(http.StatusOK, internal.NewHttpResponse(http.StatusOK, p))
	}
}
