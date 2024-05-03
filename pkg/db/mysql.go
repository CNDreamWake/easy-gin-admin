package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func initMYSQL(conf *MYSQL) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", conf.Username, conf.Password, conf.Addr, conf.Port, conf.DbName, conf.Config)
	mysqlDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 255,
	}), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		log.Fatalf("gorm open err:%v\n", err)
	}
	s, err := mysqlDB.DB()
	if err != nil {
		log.Fatalf("gorm get db err:%v\n", err)
	}
	s.SetMaxIdleConns(conf.MaxIdleConns)
	s.SetMaxOpenConns(conf.MaxOpenConns)
	log.Println("MySQL服务启动成功")
	return mysqlDB
}
