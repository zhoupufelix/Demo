package router

import (
	"github.com/julienschmidt/httprouter"
	"Demo/api/v1"
)

type Router struct {
	Method string
	Path string
	HandleFunc httprouter.Handle
}

type Routers []Router


func AllRouters()Routers{
	routers := Routers{
		Router{"GET","/v1/users/:id",v1.GetUserByUID},

	}
	return routers
}