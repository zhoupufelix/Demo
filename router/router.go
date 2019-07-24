package router

import (
	"net/http"
	"demo/controller"
	"Demo/config"
)

var mux = &http.ServeMux{}

func NewRouter()*http.ServeMux{
	//处理静态文件
	files := http.FileServer(http.Dir(config.APP_PATH))
	mux.Handle("/static/",http.StripPrefix("/static/",files))
	c := &controller.PublicController{}
	mux.HandleFunc("/",c.Login)
	return mux
}
