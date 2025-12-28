package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitesh111sinha/apis/models"
	"github.com/nitesh111sinha/apis/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(c *gin.Context) {
	username := c.Param("name")
	user, err := uc.userService.GetUser(&username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "user": user})
}

func (uc *UserController) GetAll(c *gin.Context){
	users, err := uc.userService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "users": users})	
}

func (uc *UserController) UpdateUser(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.userService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) ResetUser(c *gin.Context){
	username := c.Param("name")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.userService.ResetUser(&username, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	username := c.Param("name")
	if err := uc.userService.DeleteUser(&username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})	
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/user")
	router.POST("/create", uc.CreateUser)
	router.GET("/get/:name", uc.GetUser)
	router.GET("/getall", uc.GetAll)
	router.PATCH("/update", uc.UpdateUser)
	router.PUT("/reset/:name", uc.ResetUser)
	router.DELETE("/delete/:name", uc.DeleteUser)
}