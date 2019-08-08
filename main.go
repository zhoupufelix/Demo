package main

import (
	"demo/router"
	"net/http"
)

func main(){
	mux := router.NewRouter()
	s := &http.Server{
		Addr:":8089",
		Handler:mux,
	}
	s.ListenAndServe()
}
