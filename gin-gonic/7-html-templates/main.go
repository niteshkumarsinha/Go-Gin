package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// LoadHTMLGlob
	// LoadHTMLFiles
	
	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLFiles("templates/home.html", "templates/about.html", "templates/contact.html", "templates/products.html")

	r.GET("/home", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/about")
	})
	
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{
			"title": "About",
			"description": "About page",
		})
	})
	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", gin.H{
			"title": "Contact",
			"description": "Contact page",
		})
	})
	r.GET("/products", func(c *gin.Context) {
		c.HTML(http.StatusOK, "products.html", gin.H{
			"title": "Products",
			"description": "Products page",
		})
	})	


	
	log.Fatal(r.Run(address))
}
