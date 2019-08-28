package router

import (
	"github.com/julienschmidt/httprouter"
	"Demo/api/middleware/jwt"
	"net/http"
	"os"
	"github.com/swaggo/http-swagger"
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

	//静态文件处理
	r.ServeFiles("/docs/*filepath",http.Dir(os.Getenv("gopath") +"/src/Demo/api/docs/"))
	r.Handler("GET", "/swagger/*index", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8087/swagger/doc.json"), //The url pointing to API definition"
	))
	return r
}


