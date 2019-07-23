package router

import (
	"net/http"
	"demo/controller"
)

var mux = &http.ServeMux{}

func NewRouter()*http.ServeMux{
	//处理静态文件
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static",http.StripPrefix("/static/",files))
	c := &controller.PublicController{}
	mux.HandleFunc("/",c.Login)
	return mux
}
