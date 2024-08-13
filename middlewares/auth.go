package middlewares

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context)  {
	// Get Token From Header
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		// Response unauthorized and stop the server
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	// Verify the token if it is valid
	userId, err := utils.VerifyToken(token)
	 
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	// Set userId to data context
	context.Set("userId", userId)

	context.Next()
}