package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"server/pkg/cache"
	"server/pkg/db"
	logPkg "server/pkg/log"
	"server/pkg/server"
)

const filepath = "./config.toml"

var conf Config

func Initialization() *Config {
	if _, err := os.Stat(filepath); err != nil {
		log.Fatalf("未获取到配置文件: %v\n", err)
	}
	if _, err := toml.DecodeFile(filepath, &conf); err != nil {
		log.Fatalf("解析配置文件失败: %v\n", err)
	}
	log.Println("配置文件加载成功")
	return &conf
}

func GetConfig() *Config {
	return &conf
}

type Config struct {
	Server server.Server
	Db     db.DB
	Cache  cache.Cache
	Log    logPkg.Log
}
