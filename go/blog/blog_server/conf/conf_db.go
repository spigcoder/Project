package conf

import "fmt"

type DB struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Name           string `yaml:"name"`
	Debug          bool   `yaml:"debug"`
	SourceName     string `yaml:"source_name"` // mysql
	Max_idle_conns int    `yaml:"max_idle_conns"`
	Max_open_conns int    `yaml:"max_open_conns"`
}

func (db DB) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}
