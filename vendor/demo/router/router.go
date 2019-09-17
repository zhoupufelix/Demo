package router

import (
	"net/http"
	"Demo/controller"
	"Demo/controller/back"
	fp "path/filepath"
	"Demo/conf"
)

var mux = &http.ServeMux{}

func NewRouter()*http.ServeMux{
	//处理静态文件
	app_path,_:= fp.Abs(conf.Cfg.Admin.App_path)
	files := http.FileServer(http.Dir(app_path))
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

