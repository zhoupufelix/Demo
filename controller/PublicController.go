package controller

import (
	"net/http"
	"html/template"
	"Demo/config"
	"log"
	"fmt"
	"encoding/json"
	"Demo/model"
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
	var login = &model.Login{}
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form["username"][0]
		password := r.Form["password"][0]
		//登录逻辑
		rsp := login.CheckLogin(username,password)
		result["code"] = rsp.Code
		result["msg"] = rsp.Msg
	}
	jsonStr,err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w,string(jsonStr))
}

func (p *PublicController)Test(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET" {
		users :=  &model.Users{}
		lastInsertID := users.AddUsers()
		fmt.Fprint(w,lastInsertID)
	}

}

