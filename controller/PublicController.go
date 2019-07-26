package controller

import (
	"net/http"
	"html/template"
	"Demo/config"
	"log"
	"fmt"
)

type PublicController struct {

}

func (p *PublicController)Login(w http.ResponseWriter,r *http.Request){
	t, err:= template.ParseFiles(config.APP_PATH + "login.html")
	if err !=nil {
		log.Println(err)
		return
	}
	t.Execute(w,nil)
}

func (p *PublicController)DoLogin(w http.ResponseWriter,r *http.Request){
	result := map[string]interface{}{"code":0,"msg":""}
	if r.Method == "POST" {

		fmt.Fprint(w,"hello world")
	}
}



