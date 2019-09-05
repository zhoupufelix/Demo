package router

import (
	"net/http"

	"Demo/controller"
	"Demo/controller/back"
	"Demo/conf"
)

var mux = &http.ServeMux{}

func NewRouter()*http.ServeMux{
	//处理静态文件
	files := http.FileServer(http.Dir(conf.APP_PATH))
	mux.Handle("/static/",http.StripPrefix("/static/",files))

	//public
	public := &controller.PublicController{}
	mux.HandleFunc("/login/index",public.Login)
	mux.HandleFunc("/login/do",public.DoLogin)
	mux.HandleFunc("/login/test",public.Test)

	//back
	admin := &back.IndexController{}
	mux.HandleFunc("/back/index/index",admin.Index)

	return mux
}

