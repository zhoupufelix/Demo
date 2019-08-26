package v1

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/astaxie/beego/validation"
	"Demo/model"
	"Demo/pkg/e"
	"log"
	"Demo/api/libs"
	"Demo/library"
	"fmt"
	"strconv"
)


func GetAuth(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	data := make(map[string]interface{})
	valid := validation.Validation{}
	user := &model.Users{}
	user.Username = username
	user.Password = password

	ok,_ := valid.Valid(user)

	code := e.INVALID_PARAMS
	if ok {
		isExist := user.CheckUser(username, password)
		fmt.Println(isExist)
		code = e.ERROR_AUTH
		if isExist {
			token, err := library.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		}
	}else{
		for _,err := range valid.Errors{
			log.Println(err.Key,err.Message)
		}
	}

	libs.JSON(w,http.StatusOK,libs.M{
		"code":code,
		"data":data,
		"msg":e.GetMsg(code),
	})
}


func GetUserByUID(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	user := &model.Users{}
	pk := params.ByName("id")
	id,err := strconv.Atoi(pk)
	if err != nil {
		fmt.Println(err)
	}
	u,err := user.FindByPK(id)
	if err != nil {
		fmt.Println(err)
	}
	data := make(map[string]interface{})
	data["user"] = u
	code := e.SUCCESS

	libs.JSON(w,http.StatusOK,libs.M{
		"code":code,
		"data":data,
		"msg":e.GetMsg(code),
	})

}