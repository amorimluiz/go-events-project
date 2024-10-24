package controllers

import (
	"amorimluiz/events/models"
	"amorimluiz/events/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventController struct {
	service *services.EventService
}

func NewEventController() *EventController {
	return &EventController{service: services.NewEventService()}
}

func (c *EventController) CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not parse request data:\n%v", err)})
		return
	}

	userId, err := uuid.Parse(context.GetString("userId"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID."})
		return
	}

	event.UserID = userId

	context.JSON(c.service.CreateEvent(&event))
}

func (c *EventController) ListEvents(context *gin.Context) {
	context.JSON(c.service.ListEvents())
}

func (c *EventController) GetEvent(context *gin.Context) {
	id := context.Param("id")

	context.JSON(c.service.GetEvent(id))
}

func (c *EventController) UpdateEvent(context *gin.Context) {
	id, err := uuid.Parse(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID."})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not parse request data:\n%v", err)})
		return
	}

	event.ID = id

	userId, err := uuid.Parse(context.GetString("userId"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID."})
		return
	}

	context.JSON(c.service.UpdateEvent(&event, &userId))
}

func (c *EventController) DeleteEvent(context *gin.Context) {
	id := context.Param("id")

	context.JSON(c.service.DeleteEvent(id))
}
