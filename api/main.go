package main

import (
	"net/http"
	"Demo/config"
	r "Demo/api/router"
)

func main(){
	routers := r.AllRouters()
	mux := r.NewRouter(routers)
	server := http.Server{
		Addr:config.API_ADDR,
		Handler:mux,
	}
	server.ListenAndServe()
}