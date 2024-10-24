package main

import (
	"amorimluiz/events/db"
	"amorimluiz/events/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	server := gin.Default()

	router.RegisterRoutes(server)

	server.Run(":8080")
}
