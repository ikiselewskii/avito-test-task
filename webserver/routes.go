package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikiselewskii/avito-test-task/database"
	"github.com/ikiselewskii/avito-test-task/models"
)

func InitializeEndpoints(r *gin.Engine) {
	r.GET("/", rootRoute)
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/add_money", addMoney)
}

func rootRoute(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "hello world", "status": http.StatusOK})
}

func addMoney (context *gin.Context) {
	var json models.Customer
	err := context.ShouldBindJSON(&json)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "wrong input, try again"})
		return
	}
	err = database.AddMoney(json, context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "payment accepted"})
	
}
