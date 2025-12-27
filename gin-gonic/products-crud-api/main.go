package main

import (
	"database/sql"
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nitesh111sinha/products-crud-api/internal/controllers"
)

func main() {
	var r *gin.Engine = gin.Default()
	var address string = ":8080"

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error: %v", err)
	}

	r.GET("/products", controllers.GetProducts(db))
	r.GET("/products/:guid", controllers.GetProduct(db))
	r.POST("/products", controllers.PostProduct(db))
	r.PUT("/products/:guid", controllers.PutProduct(db))
	r.DELETE("/products/:guid", controllers.DeleteProduct(db))

	log.Fatal(r.Run(address))
}
