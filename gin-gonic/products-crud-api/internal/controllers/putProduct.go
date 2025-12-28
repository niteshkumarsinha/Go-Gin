package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitesh111sinha/products-crud-api/internal"
	"github.com/jinzhu/copier"
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
			c.JSON(http.StatusBadRequest, internal.NewHttpResponse(http.StatusBadRequest, err.Error()))
			return
		}

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, internal.NewHttpResponse(http.StatusBadRequest, err.Error()))
			return
		}

		ctx := c.Request.Context()

		var options = copier.Option{
			IgnoreEmpty: true,
			DeepCopy: true,
		}

		currentProduct := internal.Product{}
		copier.CopyWithOption(&currentProduct, &payload, options)
		// Perform the update
		_, err := db.ExecContext(ctx, "UPDATE products SET name = ?, price = ?, description = ? WHERE guid = ?",
			currentProduct.Name, currentProduct.Price, currentProduct.Description, binding.guid)

		if err != nil {
			c.JSON(http.StatusInternalServerError, internal.NewHttpResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		// Fetch the updated product to confirm it exists and return the new state
		var updatedProduct internal.ProductResponse
		err = db.QueryRowContext(ctx, "SELECT guid, name, price, description FROM products WHERE guid = ?", binding.guid).
			Scan(&updatedProduct.GUID, &updatedProduct.Name, &updatedProduct.Price, &updatedProduct.Description)

		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, internal.NewHttpResponse(http.StatusNotFound, "Product not found"))
				return
			}
			c.JSON(http.StatusInternalServerError, internal.NewHttpResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		fmt.Println(updatedProduct)

		c.JSON(http.StatusOK, internal.NewHttpResponse(http.StatusOK, updatedProduct))
	}
}
