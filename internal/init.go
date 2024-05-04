package internal

import (
	"github.com/gin-gonic/gin"
	"server/internal/controller"
	r "server/internal/router"
	"server/internal/service"
)

// Initialization 初始化函数，负责设置API的路由。
// group: 代表一组具有相同前缀的路由，用于定义API路由的基础路径。
func Initialization(group gin.RouterGroup) {
	// 初始化服务层，进行一些必要的服务配置。
	service.Initialization()
	// 初始化API层，注册API相关的中间件和处理函数。
	controller.Initialization()
	// 配置路由，定义API的路由结构。
	router(group)
}

// router 设置API路由的内部函数。
// group: 用于定义路由前缀的gin.RouterGroup，此参数确定了API路由的基础URL路径。
func router(group gin.RouterGroup) {
	// 创建处理前端请求的路由分组，为前端请求定义了一级路由路径。
	api := group.Group("/api")
	// 配置前端路由，注册前端相关的所有路由路径和处理函数。
	r.Api(api)
}
