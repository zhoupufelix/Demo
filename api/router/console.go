package router

import "github.com/julienschmidt/httprouter"

func NewRouter(routers Routers)(result *httprouter.Router){
	r := httprouter.New()
	if routers == nil  {
		return r
	}
	for _,router := range routers{
		var handle httprouter.Handle
		handle = router.HandleFunc

		//中间件


		r.Handle(router.Method,router.Path,handle)
	}
	return r
}


