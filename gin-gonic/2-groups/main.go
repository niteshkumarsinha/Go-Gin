package main


import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {
	var r *gin.Engine = gin.Default()
	var address string = "localhost:8080"

	v1 := r.Group("/api/v1")
	v1.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})

	var v2 = r.Group("/api/v2")
	v2.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})

	log.Fatal(r.Run(address))
}