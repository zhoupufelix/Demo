package v1

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"Demo/api"
	"strconv"
	"Demo/pkg/e"
)

func GetUserByUID(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	var rsp api.Response
	id := params.ByName("id")
	id,err := strconv.Atoi(id)
	if  err != nil || id <= 0{
		rsp.Code = e.INVALID_PARAMS
		rsp.Msg  = e.GetMsg(e.INVALID_PARAMS)
	}

}