package model

import (
	"time"
	"log"
	"Demo/library"
)


type Users struct {
	BaseModel
	ID int `db:"id" json:"id"`
	Username string `db:"username" json:"username" valid:"Required; MaxSize(50)"`
	Password string `db:"password" json:"password" valid:"Required; MaxSize(50)"`
	Salt string `db:"salt" json:"salt"`
	Create_at int `db:"create_at" json:"create_at"`
	Update_at int `db:"update_at" json:"update_at"`
	Role_id int `db:"role_id" json:"role_id"`
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

func(u *Users)CheckUser(username,password string)bool{
	user,err := u.FindByUsername(username)
	if err != nil {
		return false
	}
	passwordInput  := library.MakeMD5(password+user.Salt)
	if passwordInput != user.Password {
		return false
	}
	return true
}


func (u *Users)FindByUsername(username string)(*Users,error){
	sql := "SELECT `id`,`username`,`password`,`salt`,`create_at`,`update_at`,`role_id` FROM users WHERE username=? limit 1"
	err := u.QueryOne(sql,u,username)
	if err != nil {
		return nil,err
	}
	return u,nil
}

func(u *Users)FindByPK(pk int)(*Users,error){
	sql := "SELECT `id`,`username`,`password`,`salt`,`create_at`,`update_at`,`role_id` FROM users WHERE id=? limit 1"
	err := u.QueryOne(sql,u,pk)
	if err != nil {
		return nil,err
	}
	return u,nil
}
