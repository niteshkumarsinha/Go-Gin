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
		guid := c.Param("guid")

		result, err := db.ExecContext(c.Request.Context(), "DELETE FROM products WHERE guid = ?", guid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, internal.NewHttpResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, internal.NewHttpResponse(http.StatusNotFound, "Product not found"))
			return
		}

		c.JSON(http.StatusNoContent, internal.NewHttpResponse(http.StatusOK, "Product deleted successfully"))
	}
}
