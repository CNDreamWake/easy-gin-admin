package db

type DB struct {
	MYSQL MYSQL
}

type MYSQL struct {
	Addr         string
	Config       string
	DbName       string
	MaxIdleConns int
	MaxOpenConns int
	Port         uint
	Username     string
	Password     string
}
