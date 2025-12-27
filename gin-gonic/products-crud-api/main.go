package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/nitesh111sinha/products-crud-api/internal/controllers"
)

func main() {
	var r *gin.Engine = gin.Default()
	var address string = ":8080"

	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:guid", controllers.GetProduct)
	r.POST("/products", controllers.PostProduct)
	r.PUT("/products/:guid", controllers.PutProduct)
	r.DELETE("/products/:guid", controllers.DeleteProduct)

	log.Fatal(r.Run(address))
}
