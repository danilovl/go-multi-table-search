package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func SecurityHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		configToken := os.Getenv("X_AUTH_TOKEN")
		if len(configToken) == 0 {
			return
		}

		xAuthTokenHeader := context.Request.Header["X-Auth-Token"]
		if len(xAuthTokenHeader) == 0 {
			context.JSON(http.StatusForbidden, gin.H{"error": "Missing auth token"})
			context.Abort()

			return
		}

		xAuthToken := xAuthTokenHeader[0]
		if configToken == xAuthToken {
			return
		}

		if configToken != xAuthToken {
			context.JSON(http.StatusForbidden, gin.H{"error": "Auth token is invalid"})
			context.Abort()
		}
	}
}
