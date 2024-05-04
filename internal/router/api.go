package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/controller"
)

func Api(group *gin.RouterGroup) {
	{
		c := controller.Api
		group.GET("/test", c.Test)
	}
}
