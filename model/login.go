package model

import (
	"Demo/library"
	"log"
)

type Login struct {
	BaseModel
	response Response
}


func(l *Login)CheckLogin(username,password string)Response{
	//找出salt
	users := &Users{}
	users,err := users.FindByUsername(username)

	if err != nil {
		log.Println(err)
		l.response.Code = -1
		l.response.Msg = "用户名错误"
		return l.response
	}

	md5 := library.MakeMD5(password+users.Salt)

	if md5 != users.Password {
		l.response.Code = -2
		l.response.Msg = "密码错误"
		return l.response
	}

	return l.response
}



