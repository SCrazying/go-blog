package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type DbConfig struct {
	Type string				//数据库连接类型
	Host string				//数据库连接地址
	Port string				//数据库连接端口
}

type  HttpServerConfig struct {
	Host string
	Port string
}

type Config struct {
	DbConfig
	HttpServerConfig
}
var cfg Config

func init(){

	file ,err := os.Open("./config.yaml")
	if err!=nil{
		log.Fatalln(err)
	} else{
		yaml_data,err := ioutil.ReadAll(file)
		if err!= nil{
			log.Fatalln(err)
		}
		err = yaml.Unmarshal(yaml_data,cfg)
		if err!=nil{
			fmt.Println("加载配置文件出错 , ",err)
		}
	}
	config := &Config{
		DbConfig{
			"mysql",
			"192.168.2.1",
			"1111",
		},
		HttpServerConfig{
			"192.168.2.1",
			"11",
		},
	}
	b,err := yaml.Marshal(config)
	fmt.Println(string(b))
	fmt.Println("init")
}
