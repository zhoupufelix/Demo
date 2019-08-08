package model

import (
	"time"
)

type Users struct {
	BaseModel
}

func(u *Users)AddUsers()int64{
	sql := `INSERT INTO users(id,username,password,salt,create_at,update_at)VALUES("",?,?,?,?,?)`
	s := make([]interface{},5,5)
	s[0] = "felix"
	s[1] = "8446D39ABA030105D82A8AA8DD6B9E01"
	s[2] = "&*%FGOa2"
	s[3] = time.Now().Unix()
	s[4] = time.Now().Unix()
	lastInsertID,err := u.Insert(sql,s...)

	if err != nil {
		panic(nil)
	}
	return lastInsertID
}
