package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitesh111sinha/products-crud-api/internal"
)

// DeleteProduct deletes a product from the database
func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var guidBinding GUIDBinding

		if err := c.ShouldBindUri(&guidBinding); err != nil {
			resp := internal.NewHttpResponse(http.StatusBadRequest, err.Error())
			c.JSON(http.StatusBadRequest, resp)
			return
		}

		guid := guidBinding.GUID
		result, e := db.ExecContext(c.Request.Context(), "DELETE FROM products WHERE guid = ?", guid)
		if e != nil {
			resp := internal.NewHttpResponse(http.StatusInternalServerError, e.Error())
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		rowsAffected, e := result.RowsAffected()
		if e != nil {
			resp := internal.NewHttpResponse(http.StatusInternalServerError, e.Error())
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		if rowsAffected == 0 {
			resp := internal.NewHttpResponse(http.StatusNotFound, "Product not found")
			c.JSON(http.StatusNotFound, resp)
			return
		}

		resp := internal.NewHttpResponse(http.StatusOK, "Product deleted successfully")
		c.JSON(http.StatusNoContent, resp)
	}
}
