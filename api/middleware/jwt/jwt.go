package jwt

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"Demo/pkg/e"
)


func JWT(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params){
		//jwt判断
		var code int
		var data interface{}

		token := param.ByName("token")
		if token == "" {
			code = e.INVALID_PARAMS

		}

		fn(w,r,param)
	}
}
