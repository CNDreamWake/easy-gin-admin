package cache

type Cache struct {
	Redis Redis
}

type Redis struct {
	Addr     string
	Db       int
	Password string
}
