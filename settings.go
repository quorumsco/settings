package settings

import (
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

type TOMLConfig struct {
	Debug    bool
	Migrate  bool
	Server   Server
	Database map[string]interface{}
}

var Default = TOMLConfig{
	Debug:   true,
	Migrate: false,
	Server: Server{
		Host: "0.0.0.0",
		Port: 8080,
	},
	Database: map[string]interface{}{
		"Client": "sqlite3",
		"Path":   DefaultSqlite3.Path,
	},
}

func (config TOMLConfig) Client() string {
	client, ok := config.Database["client"].(string)
	if client != "" && ok {
		return client
	}
	return Default.Database["client"].(string)
}

func (config TOMLConfig) DB() (dialect, args string) {
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
