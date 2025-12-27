package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

var accounts = map[string]string{
	"admin": "admin",
	"user":  "user",
}

func main() {
	var r *gin.Engine = gin.Default()
	var address string = ":8080"


	r.Use(func(ctx *gin.Context) {
		var requestID = ctx.GetHeader("X-Request-Id")
		if len(requestID) == 0 {
			id := uuid.New().String()
			ctx.Writer.Header().Set("X-Request-Id", id)
		} else {
			ctx.Writer.Header().Set("X-Request-Id", requestID)
		}

		fmt.Println("logging...")
		fmt.Println("X-Request-Id: ", ctx.Writer.Header().Get("X-Request-Id"))
		ctx.Next()
	})
	// For every request
	// r.Use(gin.BasicAuth(accounts))

	var authMiddleware = gin.BasicAuth(accounts)

	r.GET("/ping", authMiddleware, func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/customer", func(c *gin.Context) {
		var customer = Customer{
			Email: "test@gmail.com",
			Role: "BASIC",
			StreetAddress: "123 Main St",
			StreetNumber: "123",
		}	
		c.JSON(http.StatusOK, customer)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	log.Fatal(r.Run(address))
}
