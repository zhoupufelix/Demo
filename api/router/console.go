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

		handle = jwt.JWT(handle)

		r.Handle(router.Method,router.Path,handle)
	}
	return r
}


