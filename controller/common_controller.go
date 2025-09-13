package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matt0792/mscommon/microservice"
)

type CommonController struct{}

func NewCommonController() microservice.RouteRegistrar {
	return &CommonController{}
}

func (ctrl *CommonController) RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", ctrl.Ping)
}

func (ctrl *CommonController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
