package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type dBConfig struct {
	Type string				//数据库连接类型
	Host string				//数据库连接地址
	Port string				//数据库连接端口
	UserName string			//数据库连接名
	Password string 		//数据库密码
	Charset string 			//数据库字体
	Database string			//数据库名
	Sslmode string			//ssl
	Url string 				//数据库
}
type serverConfig struct {
	Host string
	Port string
	Url string
}

type Config struct {
	Db dBConfig
	Server serverConfig
}

func (c *Config)GetUrl()(string,string){
	switch c.Db.Type {
	case "mysql":
		c.Db.Url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			c.Db.UserName,
			c.Db.Password,
			c.Db.Host,
			c.Db.Port,
			c.Db.Database,
			c.Db.Charset)
		break
	case "postgresql":
		c.Db.Url = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			c.Db.Host,
			c.Db.Port,
			c.Db.UserName,
			c.Db.Database,
			c.Db.Password,
			c.Db.Sslmode)
		break
	}

	return c.Db.Type,c.Db.Url
}
func (c *Config) GetServer()string{
	c.Server.Url = fmt.Sprintf("%s:%s",c.Server.Host,c.Server.Port)
	return c.Server.Url
}

var  Cfg = &Config{}

func init(){

	file ,err := os.Open("./config.yaml")
	if err!=nil{
		log.Fatalln(err)
	} else{
		yamlData,err := ioutil.ReadAll(file)
		if err!= nil{
			log.Fatalln(err)
		}
		err = yaml.Unmarshal(yamlData,Cfg)
		if err!=nil{
			fmt.Println("加载配置文件出错 , ",err)
		}
	}
	fmt.Println("init")
}
