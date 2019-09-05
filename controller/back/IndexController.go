package back

import (
	"net/http"
	"fmt"
	"Demo/conf"
	"log"
	"html/template"
)

type IndexController struct {
	Controller
}


func (this *IndexController) Index(w http.ResponseWriter ,r *http.Request){
	if r.Method == "GET" {
		this.template = conf.APP_PATH +"index.html"
		fmt.Println(this.template)
		t,err := template.ParseFiles(this.template)
		if err != nil {
			log.Println(err)
			return
		}
		t.Execute(w,nil)
	}


}
