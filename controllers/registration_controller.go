package controllers

import (
	"amorimluiz/events/services"

	"github.com/gin-gonic/gin"
)

type RegistrationController struct {
	service *services.RegistrationService
}

func NewRegistrationController() *RegistrationController {
	return &RegistrationController{service: services.NewRegistrationService()}
}

func (c *RegistrationController) Register(context *gin.Context) {
	userID := context.GetString("userId")
	eventID := context.Param("id")

	context.JSON(c.service.RegisterUserInEvent(userID, eventID))
}

func (c *RegistrationController) Unregister(context *gin.Context) {
	userID := context.GetString("userId")
	eventID := context.Param("id")

	context.JSON(c.service.UnregisterUserFromEvent(userID, eventID))
}
