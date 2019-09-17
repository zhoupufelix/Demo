package main

import (
	"demo/router"
	"net/http"
	"Demo/library/session"
	"Demo/conf"
)

var GlobalSessions *session.Manager

//init函数中初始化
func init() {
	GlobalSessions, _ = session.NewManager("file", "gosessionid", 3600,"runtime")
	go GlobalSessions.SessionGC()
}

func main(){
	mux := router.NewRouter()
	s := &http.Server{
		Addr:conf.Cfg.Admin.Server+conf.Cfg.Admin.Port,
		Handler:mux,
	}
	s.ListenAndServe()
}
