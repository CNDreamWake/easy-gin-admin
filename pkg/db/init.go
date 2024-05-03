package db

import "gorm.io/gorm"

var db *gorm.DB

func Initialization(conf *DB) {
	db = initMYSQL(&conf.MYSQL)
}

func GetDb() *gorm.DB {
	return db
}
