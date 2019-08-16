package provider

import (
	"sync"
	"Demo/library/session"
)

var (
	filepder = &FileProvider{} //注册实例
	gcmaxlifetime int64  //最大GC时间
)

type FileProvider struct {
	lock 		sync.RWMutex
	maxlifetime int64
	savePath 	string
}


//SessionInit(sid string)(Session,error)
//SessionRead(sid string)(Session,error)
//SessionDestroy(sid string)error
//SessionGC(maxLifeTime int64)

func (fp *FileProvider)SessionInit(){

}





//register FileProvider
func init() {
	session.RegisterProvider("file", filepder)
}