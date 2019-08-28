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

// ShowUser godoc
// @tags Users
// @Summary 获得单个用户信息
// @Description get data by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param token query string true "JWT TOKEN"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 401 {string} json "{"code":400,"data":{},"msg":"请求参数错误"} {"code":20001,"data":{},"msg":"Token鉴权失败"} {"code":20002,"data":{},"msg":"Token已超时"}"
// @Router /users/{id} [get]
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
	data := u
	code := e.SUCCESS

	libs.JSON(w,http.StatusOK,libs.M{
		"code":code,
		"data":data,
		"msg":e.GetMsg(code),
	})
}


