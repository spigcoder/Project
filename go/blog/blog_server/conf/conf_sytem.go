package conf

import "fmt"

type System struct {
	IP   string `yaml:"ip"`
	PORT int `yaml:"port"`
	GinMode string `yaml:"gin_mode"`
}

func (s *System) Addr() string {
	return fmt.Sprintf("%s:%d", s.IP, s.PORT)
}
