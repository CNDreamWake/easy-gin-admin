package boot

import (
	"server/internal"
	"server/pkg/cache"
	"server/pkg/config"
	"server/pkg/db"
	"server/pkg/log"
	"server/pkg/server"
)

// Start 启动函数负责初始化并启动API服务器。
// 该函数不接受参数，也不返回任何值。
func Start() {
	// 创建一个新的服务器实例。
	app := server.NewServer()
	//加载Config配置
	conf := config.Initialization()
	//加载日志配置
	log.Initialization(&conf.Log)
	//启动数据库服务
	db.Initialization(&conf.Db)
	//启动缓存服务
	cache.Initialization(&conf.Cache)
	// 对内部系统进行初始化，配置路由等。
	internal.Initialization(app.RouterGroup)
	// 启动服务器，监听端口并处理请求。
	server.Run(app, &conf.Server)
}
