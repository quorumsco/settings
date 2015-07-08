package settings

import "fmt"

type Postgres struct {
	DB       string
	Host     string
	Port     int
	User     string
	Password string
}

func (config TOMLConfig) Postgres() Postgres {
	var pDefault = Default.Components["postgres"].(Postgres)

	pConfig, ok := config.Components["postgres"].(Postgres)
	if !ok {
		return pDefault
	}

	if pConfig.DB == "" {
		pConfig.DB = pDefault.DB
	}

	if pConfig.Host == "" {
		pConfig.Host = pDefault.Host
	}

	if pConfig.Port == 0 {
		pConfig.Port = pDefault.Port
	}

	if pConfig.User == "" {
		pConfig.User = pDefault.User
	}

	if pConfig.Password == "" {
		pConfig.Password = pDefault.Password
	}

	return pConfig
}

func (p Postgres) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.DB)
}
