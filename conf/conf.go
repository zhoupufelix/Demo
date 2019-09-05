package conf

import (
	"flag"
	"sync"
	"path/filepath"
	"log"
	"github.com/BurntSushi/toml"
)

var (
	cfg * Config
	env string
	confPath string
	once sync.Once
)


type Config struct {
	Api api
	DB database `toml:"database"`
	Admin admin
}

type api struct {
	Server string
	Port string
	Jwt_secret string
}

type database struct {
	Type string
	Server string
	Port string
	Username string
	Password string
	Dbname string
	Maxconn int
	Maxidle int
}

type admin struct {
	Server string
	Port string
}

//初始化函数 默认读取本地配置
func init(){
	flag.StringVar(&env, "env", "local", "default env")
}

func Init() *Config{
	confPath = "src/Demo/conf/local.toml"
	if env == "product" {
		confPath = "src/Demo/conf/product.toml"
	}
	once.Do(func() {
			filePath,err := filepath.Abs(confPath)
			if err != nil {
				panic(err)
			}
			log.Printf("parse toml file once.filepath:%s\n",filePath)
			if _,err = toml.DecodeFile(filePath,&cfg);err != nil {
				panic(err)
			}
		})
	return cfg
}





