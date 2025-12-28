package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nitesh111sinha/apis/controllers"
	"github.com/nitesh111sinha/apis/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
	usersService services.UserService
	usersController *controllers.UserController
	ctx context.Context
	userCollection *mongo.Collection
	mongoClient *mongo.Client
	err error
)

func init(){
	ctx = context.TODO()
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to MongoDB")
	userCollection = mongoClient.Database("usersdb").Collection("users")
	usersService = services.NewUserService(userCollection, ctx)
	usersController = controllers.NewUserController(usersService)
	server = gin.Default()
}

func main(){
	defer mongoClient.Disconnect(ctx)
	// /api/v1/user
	usersController.RegisterUserRoutes(server.Group("/api/v1"))
	log.Fatal(server.Run(":8080"))
}
