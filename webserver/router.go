package webserver

import (
	"github.com/gin-gonic/gin"
)

func CreateRouterEngine() (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	return
}
