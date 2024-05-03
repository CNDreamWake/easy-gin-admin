package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func initRedis(conf *Redis) *redis.Client {
	if conf.Addr == "" {
		return nil
	}
	cli := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           conf.Db,
		DialTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	_, err := cli.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatalf("Redis服务启动失败, err: %v\n", err)
	}
	log.Println("Redis服务启动成功")
	return cli
}
