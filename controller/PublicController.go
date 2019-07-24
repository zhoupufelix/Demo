package controller

import (
	"net/http"
	"html/template"
	"fmt"
	"Demo/config"
)

type PublicController struct {

}

func (p *PublicController)Login(w http.ResponseWriter,r *http.Request){
	t, err:= template.ParseFiles(config.APP_PATH + "login.html")
	if err !=nil {
		fmt.Println(err)
	}
	t.Execute(w,nil)
}

