package provider

import (
	"sync"
	"Demo/library/session"
	"path/filepath"
	"os"
	"time"
	"strings"
	"github.com/pkg/errors"
	"path"
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

//session 文件存储类
type FileSessionStore struct {
	sid string //session_id
	lock sync.RWMutex
	values map[string]interface{}
}

//FileSessionStore 需要实现以下方法
//Set(key,value interface{})error
//Get(key,value interface{})interface{}
//Del(key interface{})error
//SID() string

//设置session
func(fs *FileSessionStore)Set(key string,value interface{})error{
	fs.lock.Lock()
	defer fs.lock.Unlock()
	fs.values[key] = value
	return nil
}

//获取session
func(fs *FileSessionStore)Get(key string)error{
	fs.lock.Lock()
	defer fs.lock.Unlock()
	if v,ok := fs.values[key];ok {
		return v
	}
	return nil
}




//session provider 需要实现以下方法
//SessionInit(maxlifetime int64, savePath string)(Session,error)
//SessionRead(sid string)(Session,error)
//SessionDestroy(sid string)error
//SessionGC()

//自动删除过期的文件
func gcpath(path string,info os.FileInfo,err error)error{
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	if info.ModTime().Unix()+gcmaxlifetime > time.Now().Unix() {
		os.Remove(path)
	}
	return nil
}




//初始化session
func (fp *FileProvider)SessionInit(maxlifetime int64, savePath string)error{
	fp.maxlifetime = maxlifetime
	fp.savePath = savePath
	return nil
}

//读取session 如果不存在 就创建 返回session操作类
func(fp *FileProvider)SessionRead(sid string)(FileSessionStore,error){
	//如果session_id 含有./路径 直接返回空
	if strings.ContainsAny(sid, "./") {
		return nil, nil
	}
	//session_id 的长度必须大于2
	if len(sid) < 2 {
		return nil, errors.New("length of the sid is less than 2")
	}

	filepder.lock.Lock()
	defer filepder.lock.Unlock()

	err := os.MkdirAll(path.Join(fp.savePath, string(sid[0]), string(sid[1])), 0777)



}


//session 删除
func(fp *FileProvider) SessionDestoray(){
	fp.lock.Lock()
	defer fp.lock.Unlock()



}


//垃圾回收session
func (fp *FileProvider)SessionGC(){
	//上锁
	fp.lock.Lock()
	defer fp.lock.Unlock()

	gcmaxlifetime = fp.maxlifetime
	filepath.Walk(fp.savePath, gcpath)
}





//register FileProvider
func init() {
	session.RegisterProvider("file", filepder)
}