package conf

type Config struct {
	System System `yaml:"system"`
	Logs   Logs   `yaml:"log"`
	DB     DB     `yaml:"db"`  //读库
	DB1    DB     `yaml:"db1"` //写库
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
}
