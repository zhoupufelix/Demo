package router

import (
	"github.com/julienschmidt/httprouter"
	"Demo/api/middleware/jwt"
)

func NewRouter(routers Routers)(result *httprouter.Router){
	r := httprouter.New()
	if routers == nil  {
		return r
	}
	for _,router := range routers{
		var handle httprouter.Handle
		handle = router.HandleFunc
		//获取token 的操作不需要
		if  router.Name != "GetAuth" {
			//包装一下返回
			handle = jwt.JWT(handle)
		}
		r.Handle(router.Method,router.Path,handle)
	}
	return r
}


