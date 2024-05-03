package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// NewServer 创建并初始化一个新的gin服务器实例
// 返回值 *gin.Engine 返回一个配置好的gin引擎实例，可用于处理HTTP请求
func NewServer() *gin.Engine {
	app := gin.New() // 创建一个新的gin引擎
	// 使用gin的Recovery中间件，用于在panic时恢复
	app.Use(gin.Recovery())
	// 使用gin的Logger中间件，用于日志记录
	app.Use(gin.Logger())
	// 使用自定义的Cors中间件，用于处理跨域请求
	app.Use(cors)
	return app // 返回配置好的gin引擎实例
}

func Run(app *gin.Engine, conf *Server) {
	gin.SetMode(conf.Mode)
	go func() {
		app.Run(fmt.Sprintf(":%d", conf.Port))
	}()
	// 接收到退出命令后做一些操作（kill -9 强杀命令不起作用)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//处理退出的逻辑
	fmt.Println("关闭server")
}

// cors 是一个自定义的中间件，用于处理跨域请求
// 参数 c *gin.Context 代表当前的HTTP请求上下文
func cors(c *gin.Context) {
	method := c.Request.Method               // 获取当前请求的方法
	origin := c.Request.Header.Get("Origin") // 获取请求头中的Origin字段，用于判断是否跨域

	// 如果是跨域请求，则设置相应的Access-Control-*头部
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", "*")                                                                                                                          // 允许所有来源访问
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")                                                                                   // 允许的请求方法
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Accept-URL")                                                 // 允许的请求头部
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type") // 允许响应中暴露的头部
		c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                  // 允许使用凭证（如cookies）
	}

	// 如果请求方法为OPTIONS，则直接返回200，不继续处理后续中间件
	if strings.ToUpper(method) == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
	}

	c.Next() // 继续处理后续的中间件和处理函数
}
