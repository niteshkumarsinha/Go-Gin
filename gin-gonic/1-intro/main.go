package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	address := "localhost:8080"

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to my app!")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})

	r.Handle(http.MethodGet, "/welcome", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome")
	})

	r.Handle(http.MethodPost, "/post", func(c *gin.Context) {
		c.String(http.StatusOK, "Post")
	})

	log.Fatal(r.Run(address))
}