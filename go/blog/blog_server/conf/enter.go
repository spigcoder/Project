package conf

type Config struct {
	System System `yaml:"system"`
	Logs Logs `yaml:"log"`
}