package model

type Login struct {
	response Response
}


func(l *Login)CheckLogin(username,password string)Response{

	l.response.Code = -1
	l.response.Msg = "用户名错误"
	return l.response
}