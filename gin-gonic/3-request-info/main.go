package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	var r *gin.Engine = gin.Default()
	var address string = ":8080"

	// c.Param
	// c.Query
	// c.DefaultQuery
	// c.PostForm
	// c.DefaultPostForm
	// c.GetHeader
	// c.GetRawData
	// c.Get

	r.GET("/hello/:id", func(c *gin.Context) {
		var id = c.Param("id")
		fmt.Println("id: ", id)

		var name = c.DefaultQuery("name", "")
		fmt.Println("name: ", name)

		var age = c.DefaultQuery("age", "18")
		fmt.Println("age: ", age)

		c.String(http.StatusOK, "Hello %s", id)
	})

	r.POST("/products/:id", func(c *gin.Context) {
		
		var id = c.PostForm("id")
		fmt.Println("id: ", id)

		var name = c.DefaultPostForm("name", "")
		fmt.Println("name: ", name)

		var age = c.DefaultPostForm("age", "18")
		fmt.Println("age: ", age)

		c.String(http.StatusOK, "Hello %s", id)
	})

	log.Fatal(r.Run(address))
}