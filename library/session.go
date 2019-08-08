package library

import (
	"sync"
	"io"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"time"
)

type Manager struct {
	cookieName 	string
	lock 		sync.Mutex
	provider 	Provider
	maxLifeTime int64
}

//存储的接口
type Provider interface {
	SessionInit(sid string)(Session,error)
	SessionRead(sid string)(Session,error)
	SessionDestroy(sid string)error
	SessionGC(maxLifeTime int64)
}

//Session 操作接口
type Session interface {
	Set(key,value interface{})error
	Get(key,value interface{})interface{}
	Del(key interface{})error
	SID() string
}

var providers = make(map[string]Provider)

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
func(m *Manager)SessionStart(w http.ResponseWriter,r *http.Request)(Session,error){
	m.lock.Lock()
	defer m.lock.Unlock()
	cookie,err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		sid:= m.GenerateSID()
		session,_:=m.provider.SessionInit(sid)
		newCookie := http.Cookie{
			Name : m.cookieName,
			Value:session.Get(m.cookieName),
			Path:"/",
			HttpOnly:true,
			MaxAge:int(m.maxLifeTime),
		}
		http.SetCookie(w,&newCookie)
	}else{
		sid,_:=url.QueryUnescape(cookie.Value)
		session,_:= m.provider.SessionRead(sid)
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
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionGC(m.maxLifeTime)
	time.AfterFunc(time.Duration(m.maxLifeTime), func() {
		m.SessionGC()
	})
}
