package router

import (
	"net/http"

	"Demo/controller"
	"Demo/controller/back"
	"Demo/config"
	"github.com/julienschmidt/httprouter"
	"Demo/api/v1"
)

var mux = &http.ServeMux{}

func NewRouter()*http.ServeMux{
	//处理静态文件
	files := http.FileServer(http.Dir(config.APP_PATH))
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

func NewApiRouter()*httprouter.Router{
	mux := httprouter.New()
	mux.GET("/v1/users/:id",v1.GetUserByUID)

	return mux
}

