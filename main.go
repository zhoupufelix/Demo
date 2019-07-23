package main

import (
	"demo/router"
	"net/http"
)

func main(){
	mux := router.NewRouter()
	s := &http.Server{
		Addr:"0.0.0.0:8089",
		Handler:mux,
	}
	s.ListenAndServe()
}
