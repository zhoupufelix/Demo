package model

import (
	"database/sql"
	"fmt"
	"Demo/conf"
	"log"
	"reflect"
	_ "github.com/go-sql-driver/mysql"
)

type BaseModel struct {

}


var (
	db *sql.DB //声明DB的结构体对象
	err error
	tx *sql.Tx
)

type Response struct {
	Code int
	Msg string
	Data interface{}
}

func init(){
	db ,err = sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&allowOldPasswords=1",
		conf.MYSQL_USERNAME,
		conf.MYSQL_PASSWORD,
		conf.MYSQL_HOST,
		conf.MYSQL_DBNAME,
		))
	if err != nil {
		log.Fatal(err)
	}
	//设置最大空闲连接数
	db.SetMaxIdleConns(conf.MYSQL_MAXIDLE)
	//设置最大连接数
	db.SetMaxOpenConns(conf.MYSQL_MAXCONNS)
}


//insret data by sql
func (*BaseModel)Insert(sql string,args ...interface{})(lastInsertId int64, err error){
	stmt,err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	result,err := stmt.Exec(args...)
	if err != nil {
		return
	}
	lastInsertId,err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

//delete data by sql
func (*BaseModel)Delete(sql string,args ...interface{})(effectRows int64,err error){
	stmt,err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	result,err := stmt.Exec(args...)
	if err != nil {
		return
	}
	effectRows,err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}

//update data by sql
func (*BaseModel)Update(sql string,args ... interface{})(effectRows int64,err error){
	stmt,err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	result,err := stmt.Exec(args...)
	if err != nil {
		return
	}
	effectRows,err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}

func (*BaseModel)QueryAll(sql string,struc interface{},args ...interface{})(*[]interface{},error){
	stmt,err := db.Prepare(sql)
	if err!= nil {
		return nil,err
	}
	defer stmt.Close()

	rows,err := stmt.Query(args...)
	if err!= nil {
		return nil,err
	}
	defer rows.Close()
	slice := make([]interface{},0)
	s := reflect.ValueOf(struc).Elem()
	length := s.NumField()
	onerow := make([]interface{},length)
	for i:=0;i<length ;i++  {
		onerow[i] = s.Field(i).Addr().Interface()
	}

	for rows.Next(){
		err = rows.Scan(onerow...)
		if err != nil {
			log.Println(err)
			return nil,err
		}
		slice = append(slice,s.Interface())
	}

	return &slice,nil
}

func(*BaseModel)QueryOne(sql string,struc interface{},args ...interface{})(err error){
	stmt,err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(args...)

	s := reflect.ValueOf(struc).Elem()
	length := s.NumField()
	out := make([]interface{},0)

	for i:=0;i<length ;i++  {
		//去除类型是对象的
		if  s.Field(i).Type().Kind() == reflect.Struct {
			continue
		}

		out = append(out,s.Field(i).Addr().Interface())
	}

	err = row.Scan(out...)
	if err != nil {
		return
	}
	return
}




func (*BaseModel)CloseDB(){
	defer db.Close()
}

func (*BaseModel)Ping(){
	defer db.Ping()
}


/***********************************************************事务相关************************************************************/
func (*BaseModel)Begin()(tx *sql.Tx, err error){
	tx,err = db.Begin()
	return tx,err
}

func (*BaseModel)Rollback() error{
	err = tx.Rollback()
	return err
}

func (*BaseModel)Commit()error{
	err = tx.Commit()
	return err
}

func (*BaseModel)TransactionExec(sql string,args ...interface{})(effectRows int64,err error){
	stmt,err := tx.Prepare(sql)
	if err != nil {
		log.Printf("Transaction:Prepare sql failed.err=%v",err)
		return
	}
	defer stmt.Close()
	result,err := stmt.Exec(args)
	if err != nil{
		log.Printf("Transaction:Exec sql failed.err=%v",err)
		return
	}
	effectRows,err = result.RowsAffected()
	if err != nil {
		log.Printf("Transaction:Get EffectRows failed.err=%v",err)
		return
	}
	return effectRows,err
}
