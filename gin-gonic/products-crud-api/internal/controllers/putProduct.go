package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/nitesh111sinha/products-crud-api/internal"
)

type guidBinding struct {
	guid string `uri:"guid" binding:"required"`
}

type PutProductPayload struct {
	Name        string  `json:"name" binding:"required_without_all=price description"`
	Price       float64 `json:"price" binding:"required_without_all=name description,gt=0"`
	Description string  `json:"description" binding:"required_without_all=name price,omitempty,max=250"`
}

// PutProduct updates an existing product
func PutProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var binding guidBinding
		var payload PutProductPayload

		if err := c.ShouldBindUri(&binding); err != nil {
			resp := internal.NewHttpResponse(http.StatusBadRequest, err.Error())
			c.JSON(http.StatusBadRequest, resp)
			return
		}

		if err := c.ShouldBindJSON(&payload); err != nil {
			resp := internal.NewHttpResponse(http.StatusBadRequest, err.Error())
			c.JSON(http.StatusBadRequest, resp)
			return
		}

		var ctx = c.Request.Context()
		var row = db.QueryRowContext(ctx, "SELECT name, price, description FROM products WHERE guid=?", binding.guid)
		var currentProduct internal.Product

		if err := row.Scan(&currentProduct.Name, &currentProduct.Price, &currentProduct.Description); err != nil {
			if err == sql.ErrNoRows {
				resp := internal.NewHttpResponse(http.StatusNotFound, "Product not found")
				c.JSON(http.StatusNotFound, resp)
				return
			}
			resp := internal.NewHttpResponse(http.StatusInternalServerError, err.Error())
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		var option copier.Option = copier.Option{
			IgnoreEmpty: true,
			DeepCopy: true,
		}

		if e := copier.CopyWithOption(&currentProduct, &payload, option); e != nil {
			resp := internal.NewHttpResponse(http.StatusInternalServerError, e.Error())
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		_, err := db.ExecContext(ctx, "UPDATE products SET name = ?, price = ?, description = ? WHERE guid = ?", currentProduct.Name, currentProduct.Price, currentProduct.Description, binding.guid)

		if err != nil {
			resp := internal.NewHttpResponse(http.StatusInternalServerError, err.Error())
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		var updatedRow = db.QueryRowContext(ctx, "SELECT name, price, description FROM products WHERE guid=?", binding.guid)
		var product internal.ProductResponse
		if err := updatedRow.Scan(&product.Name, &product.Price, &product.Description); err != nil {
			if err == sql.ErrNoRows {
				resp := internal.NewHttpResponse(http.StatusNotFound, "Product not found")
				c.JSON(http.StatusNotFound, resp)
				return
			}
			resp := internal.NewHttpResponse(http.StatusInternalServerError, err.Error())
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		resp := internal.NewHttpResponse(http.StatusOK, product)

		c.JSON(http.StatusOK, resp)
	}
}
