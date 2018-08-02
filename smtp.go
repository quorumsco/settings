package settings

import (
	"encoding/json"
	"fmt"

	"github.com/quorumsco/logs"
)

type Smtp struct {
	Host       string
	Smtpserver string
	Port       string
	User       string
	Password   string
}

func (config Config) Smtp() (Smtp, error) {
	var smtp Smtp

	err := json.Unmarshal(config.Components["smtp"], &smtp)
	if err != nil {
		logs.Warning("%s: %s", err.Error(), "missing or wrong 'smtp' configuration, ignoring")
	}

	if smtp.Host == "" {
		smtp.Host = "https://api.quorumapps.com"
		logs.Warning("missing smtp 'host' configuration, assuming default value: 'https://api.quorumapps.com'")
	}
	if smtp.Port == "" {
		smtp.Port = "587"
		logs.Warning("missing smtp 'port' configuration, assuming default value: '587'")
	}
	if smtp.Smtpserver == "" {
		smtp.Smtpserver = "smtp.gmail.com"
		logs.Warning("missing smtp 'Smtpserver' configuration, assuming default value: 'smtp.gmail.com'")
	}
	if smtp.User == "" {
		smtp.User = "team@quorumapp.co"
		logs.Warning("missing smtp 'user' configuration, assuming default value: 'team@quorumapp.co'")
	}

	if smtp.Password == "" {
		logs.Warning("missing smtp 'password' configuration, Default password set")
		smtp.Password = "pierrotlefou"
	}

	return smtp, nil
}

func (p Smtp) String() string {
	return fmt.Sprintf("%s:%s:%s:%s",
		p.User, p.Password, p.Host, p.Port)
}
