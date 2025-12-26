package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UriBinding struct {
	ID string `uri:"id"`
}

type QueryBinding struct {
	Name string `query:"name"`
	Age string `query:"age"`
}

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
}

type ProductForm struct {
	ID string `form:"id"`
	Name string `form:"name"`
	Price string `form:"price"`
}

type HeaderBinding struct {
	RequestId string `header:"X-Request-Id"`
}

func main() {
	var r *gin.Engine = gin.Default()
	var address string = ":8080"

	// c.ShouldBind
	// c.ShouldBindQuery
	// c.ShouldBindUri
	// c.ShouldBindJSON
	// c.ShouldBindXML
	// c.ShouldBindYAML
	// c.ShouldBindQuery

	// URI binding
	r.POST("/products/:id", func(c *gin.Context) {
		var binding UriBinding
		var headerBinding HeaderBinding

		if err := c.ShouldBindHeader(&headerBinding); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		fmt.Println("HeaderBinding: ", headerBinding)

		if err := c.ShouldBindUri(&binding); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		fmt.Println("Binding: ", binding)
		c.String(http.StatusOK, "HeaderBinding: %s, id: %s", headerBinding.RequestId, binding.ID)	
	})

	// Form binding
	r.POST("/products/create", func(c *gin.Context) {
		var binding ProductForm
		if err := c.ShouldBind(&binding); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		fmt.Println("Binding: ", binding)
		c.String(http.StatusOK, "id: %s, name: %s, price: %s", binding.ID, binding.Name, binding.Price)	
	})

	// Query binding
	r.GET("/products", func(c *gin.Context) {
		var binding QueryBinding
		if err := c.ShouldBindQuery(&binding); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		fmt.Println("Binding: ", binding)
		c.String(http.StatusOK, "name: %s, age: %s", binding.Name, binding.Age)	
	})

	// JSON binding
	r.POST("/products", func(c *gin.Context) {
		var binding Product
		if err := c.ShouldBindJSON(&binding); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		fmt.Println("Binding: ", binding)
		c.String(http.StatusOK, "id: %s, name: %s, price: %s", binding.ID, binding.Name, binding.Price)	
	})

	log.Fatal(r.Run(address))
}