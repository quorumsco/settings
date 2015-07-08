package settings

import "fmt"

type Server struct {
	Host string
	Port int
}

func (config TOMLConfig) Server() Server {
	return config.Components["server"].(Server)
}

func (s Server) String() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
