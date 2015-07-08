package settings

import (
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

type TOMLConfig struct {
	Components map[string]interface{}
	Settings   map[string]interface{}
}

var Default = TOMLConfig{
	Components: map[string]interface{}{
		"server": Server{
			Host: "0.0.0.0",
			Port: 8080,
		},
		"postgres": Postgres{
			User:     "postgres",
			Password: "postgres",
			Host:     "0.0.0.0",
			Port:     5432,
			DB:       "postgres",
		},
		"sqlite3": Sqlite3{
			Path: "/tmp/users.sqlite",
		},
		"redis": Redis{
			Host: "redis",
			Port: 6379,
		},
	},
	Settings: map[string]interface{}{
		"debug":   true,
		"migrate": false,
		"client":  "sqlite3",
	},
}

func (config TOMLConfig) Client() string {
	client, ok := config.Settings["client"].(string)
	if client != "" && ok {
		return client
	}
	return Default.Settings["client"].(string)
}

func (config TOMLConfig) SqlDB() (dialect, args string) {
	switch config.Client() {
	case "postgres":
		dialect = "postgres"
		args = config.Postgres().String()
	case "sqlite3":
		fallthrough
	default:
		dialect = "sqlite3"
		args = config.Sqlite3().String()
	}
	return
}

func Parse(file string) (TOMLConfig, error) {
	f, err := os.Open(file)
	if err != nil {
		return Default, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return Default, err
	}

	var config TOMLConfig
	if err := toml.Unmarshal(buf, &config); err != nil {
		return Default, err
	}

	return config, nil
}
