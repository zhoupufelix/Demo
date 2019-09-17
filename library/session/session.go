package session

import (
	"sync"
	"io"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"time"
	"log"
	"os"
	"bytes"
	"encoding/gob"
	"fmt"
)

type Manager struct {
	cookieName 	string
	savePath 	string
	lock 		sync.Mutex
	provider 	Provider
	maxLifeTime int64
}

//providers 存放了各种session存储介质
var providers = make(map[string]Provider)

func NewManager(providerName ,cookieName string,maxLifeTime int64,savePath string) (*Manager, error){
	provider,ok := providers[providerName]
	if !ok  {
		return nil,fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}
	return &Manager{cookieName:cookieName ,provider:provider,maxLifeTime:maxLifeTime,savePath:savePath},nil
}


//session管理的接口
type Provider interface {
	SessionInit(maxlifetime int64, savePath string)error
	SessionRead(sid string)(Store,error)
	SessionDestroy(sid string)error
	SessionGC()
}

//Session 存储操作接口
type Store interface {
	Set(key,value interface{})error
	Get(key interface{})interface{}
	Del(key interface{})error
	SID() string
}



// SLogger 定制化session 日志
var SLogger = NewSessionLog(os.Stderr)


//注册Session 寄存器
func RegisterProvider(name string,provider Provider){
	if provider != nil {
		panic("session:Register provider is nil")
	}
	if _,ok:= providers[name];ok {
		panic("session:Register provider is existed")
	}
	providers[name] = provider
}

//生成唯一的Session ID
func(m *Manager)GenerateSID()string{
	b:= make([]byte,32)
	if _,err:=io.ReadFull(rand.Reader,b);err!=nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//SessionStart 开启 session
func(m *Manager)SessionStart(w http.ResponseWriter,r *http.Request)(session Store){
	m.lock.Lock()
	defer m.lock.Unlock()
	cookie,err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		sid:= m.GenerateSID()
		err =m.provider.SessionInit(m.maxLifeTime,m.savePath)
		newCookie := http.Cookie{
			Name : m.cookieName,
			Value: url.QueryEscape(sid),
			Path:"/",
			HttpOnly:true,
			MaxAge:int(m.maxLifeTime),
		}
		http.SetCookie(w,&newCookie)
	}else{
		sid,_:=url.QueryUnescape(cookie.Value)
		session,_ = m.provider.SessionRead(sid)
	}
	return
}

//SessionDestory 注销Session
func(m *Manager)SessionDestory(w http.ResponseWriter,r *http.Request){
	cookie,err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionDestroy(cookie.Value)
	expiredTime := time.Now()
	newCookie := http.Cookie{
		Name:m.cookieName,
		Path:"/",
		HttpOnly:true,
		Expires:expiredTime,
		MaxAge:-1,
	}
	http.SetCookie(w,&newCookie)
}


//SessionGC Session 垃圾回收
func(m *Manager)SessionGC(){
	//m.lock.Lock()
	//defer m.lock.Unlock()
	//m.provider.SessionGC()

	//time.AfterFunc(time.Duration(m.maxLifeTime), func() {
	//	m.SessionGC()
	//})
}

// Log implement the log.Logger
type Log struct {
	*log.Logger
}

func NewSessionLog(out io.Writer)*Log{
	sl := new(Log)
	sl.Logger = log.New(out, "[SESSION]", 1e9)
	return sl
}

// DecodeGob decode data to map
func DecodeGob(encoded []byte) (map[interface{}]interface{}, error) {
	buf := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(buf)
	var out map[interface{}]interface{}
	err := dec.Decode(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}