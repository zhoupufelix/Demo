package conf

import (
	"sync"
	"path/filepath"
	"log"
	"github.com/BurntSushi/toml"
)

var (
	Cfg Config
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
	App_path string
}

//初始化函数 默认读取本地配置
//func init(){
//	flag.StringVar(&env, "env", "local", "default env")
//}

func init(){
	confPath = "src/Demo/conf/local.toml"
	//if env == "product" {
	//	confPath = "src/Demo/conf/product.toml"
	//}
	once.Do(func() {
			filePath,err := filepath.Abs(confPath)
			if err != nil {
				panic(err)
			}
			log.Printf("parse toml file once.filepath:%s\n",filePath)
			if _,err = toml.DecodeFile(filePath,&Cfg);err != nil {
				panic(err)
			}
		})
}






