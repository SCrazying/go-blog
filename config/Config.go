package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type DbConfig struct {
	Type string //数据库连接类型
	Host string //数据库连接地址
	Port int    //数据库连接端口
}

type HttpServerConfig struct {
	Host string
	Port int
}

type Config struct {
	DbConfig
	HttpServerConfig
}

func (c *Config) GetHttpAddr() string {
	return fmt.Sprintf("%s:%d", c.HttpServerConfig.Host, c.HttpServerConfig.Port)
}

var Cfg = &Config{}

func init() {

	file, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatalln(err)
	} else {
		yaml_data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("无法读取配置文件 , ", err)
			log.Fatalln(err)
		}
		err = yaml.Unmarshal(yaml_data, Cfg)
		if err != nil {
			fmt.Println("加载默认配置文件出错 , ", err)
		}
	}
	fmt.Println("加载默认配置文件成功")
}
