package webserver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikiselewskii/avito-test-task/database"
	"github.com/ikiselewskii/avito-test-task/models"
	//"github.com/ikiselewskii/avito-test-task/utils"
)

func InitializeEndpoints(r *gin.Engine) {
	r.GET("/", rootRoute)
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/get_balance", getBalance)
	r.POST("/add_money", addMoney)
	r.POST("/reserve", reserve)
}

func rootRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world", "status": http.StatusOK})
}

func addMoney(ctx *gin.Context) {
	var json models.Customer
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "wrong input, try again"})
		return
	}
	err = database.AddMoney(json, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "payment accepted"})

}
func getBalance(ctx *gin.Context) {
	var json models.Customer
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "wrong input, try again"})
		return
	}
	balance, err := database.GetBalance(json, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "wrong input, try again"})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"balance": balance})
	return
}


func reserve(ctx *gin.Context) {
	var json models.Transaction
	err := ctx.BindJSON(&json)
	if err != nil {
		log.Println("unmarshalling error ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "wrong input, try again"})
		return
	}
	json.Type = 0
	log.Println(json)
	err = database.Reserve(json, ctx)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "no user  with such id found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "insufficient balance"})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "purchase reserved succesfully"})

}
