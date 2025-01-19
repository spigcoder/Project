package conf

import "fmt"

type System struct {
	IP   string `yaml:"ip"`
	PORT int `yaml:"port"`
}

func (s *System) Addr() string {
	return fmt.Sprintf("%s:%d", s.IP, s.PORT)
}
