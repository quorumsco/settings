package settings

import "fmt"

type Redis struct {
	Host string
	Port int
}

func (r Redis) String() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

func (config TOMLConfig) Redis() Redis {
	var rDefault = Default.Components["redis"].(Redis)

	rConfig, ok := config.Components["redis"].(Redis)
	if !ok {
		return rDefault
	}

	if rConfig.Host == "" {
		rConfig.Host = rDefault.Host
	}

	if rConfig.Port == 0 {
		rConfig.Port = rDefault.Port
	}

	return rConfig
}
