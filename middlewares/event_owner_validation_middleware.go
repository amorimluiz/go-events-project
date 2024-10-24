package middlewares

import (
	"amorimluiz/events/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func IsEventOwner(context *gin.Context) {
	userId, err := uuid.Parse(context.GetString("userId"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID."})
		context.Abort()
		return
	}

	eventId := context.Param("id")

	eventRepository := repositories.NewEventRepository()

	event, err := eventRepository.FindByID(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		context.Abort()
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not the owner of this event."})
		context.Abort()
		return
	}

	context.Next()
}
