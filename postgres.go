package settings

import "fmt"

type Postgres struct {
	DB       string
	Host     string
	Port     int
	User     string
	Password string
}

var DefaultPosgres = Postgres{
	DB:       "postgres",
	Host:     "postgres",
	Port:     5432,
	User:     "postgres",
	Password: "postgres",
}

func (config TOMLConfig) Postgres() Postgres {
	var pg = DefaultPosgres

	db, ok := config.Database["db"].(string)
	if db != "" && ok {
		pg.DB = db
	}

	host, ok := config.Database["host"].(string)
	if host != "" && ok {
		pg.Host = host
	}

	port, ok := config.Database["port"].(int64)
	if ok {
		pg.Port = int(port)
	}

	user, ok := config.Database["user"].(string)
	if user != "" && ok {
		pg.User = user
	}

	password, ok := config.Database["password"].(string)
	if password != "" && ok {
		pg.Password = password
	}

	return pg
}

func (p Postgres) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.DB)
}
