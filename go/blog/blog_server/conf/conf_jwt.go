package conf

type Jwt struct {
	Secret string `yaml:"secret"`
	Issuer string `yaml:"issuer"`
	Expire int    `yaml:"expire"`	//单位小时
}