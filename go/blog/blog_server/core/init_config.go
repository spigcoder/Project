package core

import (
	"os"
	"fmt"
	"blog_server/flags"
	"blog_server/conf"
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