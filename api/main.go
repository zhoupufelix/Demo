package main

import (
	"net/http"
	"Demo/conf"
	r "Demo/api/router"
	_ "Demo/api/docs"
	"strings"
)

// @title API 文档
// @version 1.0
// @description Twusa openapi 文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.twusa.cn
// @contact.email felixfan@twusa.cn

// @host localhost:8087
// @BasePath /v1
func main(){
	routers := r.AllRouters()
	mux := r.NewRouter(routers)
	cfg := conf.Init()
	srv := []string{cfg.Api.Server,cfg.Api.Port}
	addr := strings.Join(srv,":")

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	server.ListenAndServe()
}