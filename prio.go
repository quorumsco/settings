package settings

import (
	"encoding/json"
	"fmt"
	"github.com/quorumsco/logs"
)

type Prio struct {
	Host string
	Port int
}

func (config Config) Prio() (Prio, error) {
	var prio Prio

	err := json.Unmarshal(config.Components["prio"], &prio)
	if err != nil {
		logs.Warning("%s: %s", err.Error(), "missing or wrong 'prio' configuration, ignoring")
	}

	if prio.Host == "" {
		prio.Host = "prio"
		logs.Warning("missing prio 'host' configuration, assuming default value: 'prio'")
	}

	if prio.Port == 0 {
		prio.Port = 5000
		logs.Warning("missing prio 'port' configuration, assuming default value: 5000")
	}

	return prio, nil
}

func (r Prio) String() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
