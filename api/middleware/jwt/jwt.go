package jwt

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"Demo/pkg/e"
	"Demo/api/libs"
	"Demo/library"
	"time"
)


func JWT(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params))func(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params){
		//jwt判断
		var code int
		var data interface{}
		token := r.URL.Query().Get("token")
		code = http.StatusOK
		if token == "" {
			code = e.INVALID_PARAMS
		}else{
			claims,err := library.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			libs.JSON(w,http.StatusUnauthorized,libs.M{
				"code":code,
				"msg":e.GetMsg(code),
				"data":data,
			})
			return
		}
		fn(w,r,param)
	}
}
