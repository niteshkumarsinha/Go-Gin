package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type UriBinding struct {
	ID string `uri:"id"`
}

type QueryBinding struct {
	Name string `query:"name"`
	Age  string `query:"age"`
}

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type ProductForm struct {
	ID    string `form:"id"`
	Name  string `form:"name"`
	Price string `form:"price"`
}

type HeaderBinding struct {
	RequestId string `header:"X-Request-Id"`
}

type Customer struct {
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required,password"`
	Role          string `json:"role" binding:"required,oneof=BASIC ADMIN"`
	StreetAddress string `json:"street_address" binding:"required"`
	StreetNumber  string `json:"street_number" binding:"required_with=StreetAddress"`
}

func verifyPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	var (
		hasMinLen  = len(password) >= 8
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case char == '@' || char == '$' || char == '!' || char == '%' || char == '*' || char == '#' || char == '?' || char == '&':
			hasSpecial = true
		default:
			return false
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func main() {
	var r *gin.Engine = gin.Default()
	var address string = ":8080"

	// router.Static
	// router.StaticFile
	// router.StaticFS

	r.Static("/assets", "./assets")
	r.StaticFile("/hello", "./assets/hello.txt")
	//r.StaticFS("/assets", http.Dir("./assets"))	

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", verifyPassword)
	}
	r.Use()

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

	// Custom validators
	r.POST("/customers", func(c *gin.Context) {
		var customer Customer
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Binding: ", customer)
		c.String(http.StatusOK, "email: %s, role: %s, street_address: %s, street_number: %s", customer.Email, customer.Role, customer.StreetAddress, customer.StreetNumber)
	})

	log.Fatal(r.Run(address))
}
