package webserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeEndpoints(r *gin.Engine) {
	r.GET("/", rootRoute)
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func rootRoute(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "hello world", "status": http.StatusOK})
}
