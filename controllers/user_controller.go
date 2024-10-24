package controllers

import (
	"amorimluiz/events/models"
	"amorimluiz/events/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{service: services.NewUserService()}
}

func (c *UserController) SignUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not parse request data:\n%v", err)})
		return
	}

	context.JSON(c.service.SignUp(&user))
}

func (c *UserController) Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not parse request data:\n%v", err)})
		return
	}

	context.JSON(c.service.Login(&user))
}
