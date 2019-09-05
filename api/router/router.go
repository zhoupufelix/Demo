package router

import (
	"github.com/julienschmidt/httprouter"
	"Demo/api/v1"
)

type Router struct {
	Name string
	Method string
	Path string
	HandleFunc httprouter.Handle
}

type Routers []Router



func AllRouters()Routers{
	routers := Routers{
		Router{"GetAuth","GET","/v1/auth",v1.GetAuth},
		Router{"GetUserInfo","GET","/v1/users/:id",v1.GetUserByUID},
	}

	return routers
}