package router

import (
	"amorimluiz/events/controllers"
	"amorimluiz/events/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	eventController := controllers.NewEventController()

	authenticatedRoutes := server.Group("/")
	authenticatedRoutes.Use(middlewares.Authenticate)
	authenticatedRoutes.POST("/event", eventController.CreateEvent)

	authenticatedRoutes.POST("/event/:id/register", controllers.NewRegistrationController().Register)
	authenticatedRoutes.DELETE("/event/:id/unregister", controllers.NewRegistrationController().Unregister)

	authenticatedRoutes.Use(middlewares.IsEventOwner)
	authenticatedRoutes.PUT("/event/:id", eventController.UpdateEvent)
	authenticatedRoutes.DELETE("/event/:id", eventController.DeleteEvent)

	server.GET("/event", eventController.ListEvents)
	server.GET("/event/:id", eventController.GetEvent)

	userController := controllers.NewUserController()

	server.POST("/signup", userController.SignUp)
	server.POST("/login", userController.Login)
}
