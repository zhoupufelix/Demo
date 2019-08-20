package api

import (
	"net/http"
	"encoding/json"
)

type Response struct {
	r *http.Request
	w http.ResponseWriter
}

func (rsp *Response)JSON(code int,obj interface{}){
	//返回http 状态码
	rsp.

	header := w.Header()
	header["Content-Type"] = "application/json; charset=utf-8"
	jsonBytes,err := json.Marshal()
}