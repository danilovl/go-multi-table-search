package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController struct{}

func (controller PingController) Handler() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"result": "ok"})
	}
}

func NewPingController() PingController {
	return PingController{}
}
