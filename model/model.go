package model

import (
	"database/sql"
	"fmt"
	"Demo/config"
	"log"
	"reflect"
)

var (
	db *sql.DB //声明DB的结构体对象
	err error
	tx *sql.Tx
)

func init(){
	db ,err = sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&allowOldPasswords=1",
		config.MYSQL_USERNAME,
		config.MYSQL_PASSWORD,
		config.MYSQL_HOST,
		config.MYSQL_DBNAME,
		))
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(config.MYSQL_MAXIDLE)
	db.SetMaxOpenConns(config.MYSQL_MAXCONNS)
}


//insret data by sql
func Insert(sql string,args ...interface{})(id uint, err error){
	stmt,err := db.Prepare(sql)
	if err != nil {
		log.Printf("Prepare Statement failed,err:%v",err)
		return
	}
	defer stmt.Close()
	result,err := stmt.Exec(args)
	if err != nil {
		log.Printf("Insert failed,err:%v",err)
		return
	}
	lastInsertID,err := result.LastInsertId()
	if err != nil {
		log.Printf("Get LastInsrtID failed,err:%v",err)
		return
	}
	return lastInsertID,err
}

//delete data by sql
func Delete(sql string,args ...interface{})(effectRows uint,err error){
	stmt,err := db.Prepare(sql)
	if err != nil {
		log.Printf("Prepare Statement failed,err:%v",err)
		return
	}
	defer stmt.Close()
	result,err := stmt.Exec(args)
	if err != nil {
		log.Printf("Delete failed,err:%v",err)
		return
	}
	effectRows,err := result.RowsAffected()
	if err != nil {
		log.Printf("Get EffectRows failed,err:%v",err)
		return
	}
	return effectRows,err
}

//update data by sql
func Update(sql string,args ... interface{})(effectRows uint,err error){
	stmt,err := db.Prepare(sql)
	if err != nil {
		log.Printf("Prepare Statement failed,err:%v",err)
		return
	}
	defer stmt.Close()
	result,err := stmt.Exec(args)
	if err != nil {
		log.Printf("Update failed,err:%v",err)
		return
	}
	effectRows,err := result.RowsAffected()
	if err != nil {
		log.Printf("Get EffectRows failed,err:%v",err)
		return
	}
	return effectRows,err
}

func QueryAll(sql string,struc interface{},args ...interface{})([]interface{},error){
	stmt,err := db.Prepare(sql)
	if err!= nil {
		log.Printf("Query Prepare Failed,err=%v",err)
		return
	}
	defer stmt.Close()

	rows,err := stmt.Query(args...)
	if err!= nil {
		log.Printf("Query Failed,err=%v",err)
		return
	}
	defer rows.Close()

	result := make([]interface{},0)
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
			return
		}
		result = append(result,s.Interface())
	}

	return &result
}

func QueryOne(sql string,struc interface{},args ...interface{})(*[]interface{},error){
	stmt,err := db.Prepare(sql)
	if err!= nil {
		log.Printf("Query Prepare Failed,err=%v",err)
		return
	}
	defer stmt.Close()

	result := make([]interface{},0)
	s := reflect.ValueOf(struc).Elem()
	length := s.NumField()
	onerow := make([]interface{},length)
	row := stmt.QueryRow(args...)

	err = row.Scan()



	for i:=0;i<length ;i++  {
		onerow[i] = s.Field(i).Addr().Interface()
	}



	return &result
}





func Close(){
	defer db.Close()
}

func Ping()error{
	defer db.Ping()
}


/***********************************************************事务相关************************************************************/
func Begin()(tx *sql.Tx, error){
	tx,err = db.Begin()
	return tx,err
}

func Rollback() error{
	err = tx.Rollback()
	return err
}

func Commit()error{
	err = tx.Commit()
	return err
}

func TransactionExec(sql string,args ...interface{})(effectRows uint,err error){
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
	effectRows,err := result.RowsAffected()
	if err != nil {
		log.Printf("Transaction:Get EffectRows failed.err=%v",err)
		return
	}
	return effectRows,err
}
