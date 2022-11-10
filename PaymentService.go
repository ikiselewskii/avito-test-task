package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // create router instance
	InitializeEndpoints(router)
	router.Run()

}

func InitializeEndpoints(r *gin.Engine) {
	r.GET("/", rootRoute)
	r.POST("/add_user/:name", addUser)
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func rootRoute(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "hello world", "status": http.StatusOK})
}

//func addUser(context *gin.Context) {
//	context.JSON(http.StatusOK, gin.H{"message": "user added"})
//}

func addUser(context *gin.Context) {
	name := context.Param(":name")
	context.JSON(http.StatusOK, gin.H{"username": name, "balance": 282847})
}
