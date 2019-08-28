package main

import (
	"net/http"
	"Demo/config"
	r "Demo/api/router"
	_ "Demo/api/docs"
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
	server := http.Server{
		Addr:config.API_ADDR,
		Handler:mux,
	}
	server.ListenAndServe()
}