package controller

import (
	"net/http"

	"sample-api/handler"
	"sample-api/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService handler.UserService
}


func New(userservice handler.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}


func (u *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := u.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}


func (u *UserController) GetUser(ctx *gin.Context)  {
	var username string = ctx.Param("name")
	user, err := u.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
} 


func (u *UserController) GetAll(ctx *gin.Context)  {
	users, err := u	.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}


func (u *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := u.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}


func (u *UserController) DeleteUser(ctx *gin.Context)  {
	var username string = ctx.Param("name")
	err := u.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}


func (u *UserController) RegisterUserEndpoints (rg *gin.RouterGroup) {
	userEndpoint := rg.Group("/user")
	userEndpoint.POST("/create", u.CreateUser)
	userEndpoint.GET("/get/:name", u.GetUser)
	userEndpoint.GET("/getall", u.GetAll)
	userEndpoint.PATCH("/update", u.UpdateUser)
	userEndpoint.DELETE("/delete/:name", u.DeleteUser)
}

