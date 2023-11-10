package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

var file string

func InitFlag() {
	flag.StringVar(&file, "f", "", "数据库配置文件,类型默认使用yaml")
	flag.Parse()
}

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func InitConfig(file string) {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("read config file error"))
	}
	fmt.Printf("config_file: %+v\n", viper.AllSettings())
}

func main() {
	InitFlag()
	InitConfig(file)
	fmt.Printf("config_file=%s\n", file)
}
