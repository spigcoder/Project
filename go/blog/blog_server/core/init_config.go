package core

import (
	"blog_server/conf"
	"blog_server/flags"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConf() *conf.Config {
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err)
	}
	config := new(conf.Config)
	err = yaml.Unmarshal(byteData, config)
	if err != nil {
		panic(fmt.Sprintf("yaml格式错误, err: ", err))
	}
	fmt.Println(config)
	return config
}
