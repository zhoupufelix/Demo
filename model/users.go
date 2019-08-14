package model

import (
	"time"
	"log"
)


type Users struct {
	BaseModel
	ID int `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Salt string `db:"salt"`
	Create_at int `db:"create_at"`
	Update_at int `db:"update_at"`
	Role_id int `db:"role_id"`
}


func(u *Users)AddUsers()int64{
	sql := `INSERT INTO users(id,username,password,salt,create_at,update_at)VALUES(2,?,?,?,?,?)`
	s := []interface{}{"ben","8446D39ABA030105D82A8AA8DD6B9E01","&*%FGOa2",time.Now().Unix(),time.Now().Unix()}
	lastInsertID,err := u.Insert(sql,s...)

	if err != nil {
		log.Println(err)
	}
	return lastInsertID
}



func (u *Users)FindByUsername(username string)(*Users,error){
	sql := "SELECT `id`,`username`,`password`,`salt`,`create_at`,`update_at`,`role_id` FROM users WHERE username=?"
	err := u.QueryOne(sql,u,username)
	if err != nil {
		return nil,err
	}
	return u,nil
}

