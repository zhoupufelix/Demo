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
	"io/ioutil"
	"fmt"
	"log"
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
func(fs *FileSessionStore)Set(key,value interface{})error{
	fs.lock.Lock()
	defer fs.lock.Unlock()
	fs.values[key] = value
	return nil
}

//获取session
func(fs *FileSessionStore)Get(key interface{})interface{}{
	fs.lock.Lock()
	defer fs.lock.Unlock()
	if v,ok := fs.values[key];ok {
		return v
	}
	return nil
}

// Delete value in file session by given key
func (fs *FileSessionStore) Del(key interface{}) error {
	fs.lock.Lock()
	defer fs.lock.Unlock()
	delete(fs.values, key)
	return nil
}


// SessionID Get file session store id
func (fs *FileSessionStore) SID() string {
	return fs.sid
}



//session provider 需要实现以下方法
//SessionInit(maxlifetime int64, savePath string)(Session,error)
//SessionRead(sid string)(Session,error)
//SessionDestroy(sid string)error
//SessionGC()






//初始化session
func (fp *FileProvider)SessionInit(maxlifetime int64, savePath string)error{
	fp.maxlifetime = maxlifetime
	fp.savePath = savePath
	return nil
}

//读取session 如果不存在 就创建 返回session操作类
func(fp *FileProvider)SessionRead(sid string)(session.Store,error){
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

	//创建session存放的文件夹
	err := os.MkdirAll(path.Join(fp.savePath, string(sid[0]), string(sid[1])), 0777)
	if err != nil {
		session.SLogger.Println(err.Error())
	}

	//返回文件信息，如果有错误返回patherror
	_, err = os.Stat(path.Join(fp.savePath, string(sid[0]), string(sid[1]), sid))
	var f *os.File

	if err != nil {
		return nil, err
	}

	if os.IsNotExist(err) {
		f, err = os.Create(path.Join(fp.savePath, string(sid[0]), string(sid[1]), sid))
	}

	//打开文件句柄
	f, err = os.OpenFile(path.Join(fp.savePath, string(sid[0]), string(sid[1]), sid), os.O_RDWR, 0777)
	defer f.Close()

	os.Chtimes(path.Join(fp.savePath, string(sid[0]), string(sid[1]), sid), time.Now(), time.Now())
	var kv map[interface{}]interface{}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		kv = make(map[interface{}]interface{})
	} else {
		kv, err = session.DecodeGob(b)
		if err != nil {
			return nil, err
		}
	}

	ss := &FileSessionStore{sid: sid, values: kv}
	return ss, nil
}


//session 删除
func(fp *FileProvider) SessionDestroy(sid string)error{
	fp.lock.Lock()
	defer fp.lock.Unlock()
	os.Remove(path.Join(fp.savePath, string(sid[0]), string(sid[1]), sid))
	return nil
}


//垃圾回收session
func (fp *FileProvider)SessionGC(){
	//上锁
	fp.lock.Lock()
	defer fp.lock.Unlock()
	gcmaxlifetime = fp.maxlifetime
	filepath.Walk(fp.savePath, gcpath)
}

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

//register FileProvider
func init() {
	session.RegisterProvider("file", filepder)
}