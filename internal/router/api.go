package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/api"
)

func FrontendRouter(group *gin.RouterGroup) {
	{
		c := api.Api
		group.GET("/test", c.Test)
	}
}
