package main

import (
	"Demo/router"
	"net/http"
	"Demo/config"
)

func main(){
	mux := router.NewApiRouter()
	server := http.Server{
		Addr:config.API_ADDR,
		Handler:mux,
	}
	server.ListenAndServe()
}