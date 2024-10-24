package middlewares

import (
	"amorimluiz/events/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	parts := strings.Split(context.GetHeader("Authorization"), "Bearer ")

	if len(parts) != 2 {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid auth token"})
		context.Abort()
		return
	}

	tokenString := parts[1]
	authService := services.NewAuthService()
	token, err := authService.ValidateToken(tokenString)

	if err != nil || !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid auth token"})
		context.Abort()
		return
	}

	userId, err := token.Claims.GetSubject()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid auth token"})
		context.Abort()
		return
	}

	context.Set("userId", userId)

	context.Next()
}
