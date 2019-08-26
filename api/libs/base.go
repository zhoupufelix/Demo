package libs

import (
	"net/http"
	"encoding/json"
)

//short-cut for map
type M map[string]interface{}


func JSON(w http.ResponseWriter,code int,obj interface{}){
	//返回http 状态码
	w.WriteHeader(code)

	header := w.Header()
	header["Content-Type"] = []string{"application/json; charset=utf-8"}
	jsonBytes,err := json.MarshalIndent(obj,"","\t")
	if err != nil {
		panic(err)
	}
	w.Write(jsonBytes)
}