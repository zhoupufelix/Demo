package v1

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"Demo/api"
	"strconv"
)

func GetUserByUID(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	var rsp api.Response
	id := params.ByName("id")
	id,err := strconv.Atoi(id)
	if  err != nil {
		rsp.Code = -1
		rsp.Msg  = "参数类型错误"
	}

}