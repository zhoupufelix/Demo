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


//register FileProvider
func init() {
	session.RegisterProvider("file", filepder)
}