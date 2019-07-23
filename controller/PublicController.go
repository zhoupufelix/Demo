package controller

import (
	"net/http"
	"html/template"
)

type PublicController struct {

}

func (p *PublicController)Login(w http.ResponseWriter,r *http.Request){
	t, _ := template.ParseFiles("/src/demo/public/login.html")  	//解析模板文件
	t.Execute(w,"hello")  //执行模板的merger操作
}

