package main

import (
    "fmt"
    "context"
    "log"

    "sample-api/controller"
    "sample-api/handler"

    "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	userService          handler.UserService
	userController          controller.UserController
	ctx         context.Context
	userCollection       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userCollection = mongoclient.Database("sample-api").Collection("users")
	userService = handler.NewUserService(userCollection, ctx)
	userController = controller.New(userService)
	server = gin.Default()
}


func main() {
    defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	userController.RegisterUserEndpoints(basepath)

	log.Fatal(server.Run(":9090"))
}