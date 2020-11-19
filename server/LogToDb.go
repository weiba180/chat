package server

import (
	"database/sql"
	"fmt"
	"time"

	//初始化
	_ "github.com/go-sql-driver/mysql"
)

var MysqlDb *sql.DB
var dbErr error

const (
	USER_NAME = "mask"
	PASS_WORD = "123qwe"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "chat"
	CHARSET   = "UTF8"
)

func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	MysqlDb, dbErr = sql.Open("mysql", dbDSN)
	if dbErr != nil {
		fmt.Println("Mysql Open:", dbErr)
	}
	MysqlDb.SetMaxOpenConns(10)
	MysqlDb.SetMaxIdleConns(50)
	MysqlDb.SetConnMaxLifetime(30 * time.Second)
	//Ping确定连接是否可用
	if dbErr = MysqlDb.Ping(); dbErr != nil {
		panic("Mysql 数据库连接失败：" + dbErr.Error())
	}
}

//LogToDb 插入数据到数据库
func LogToDb(msg string, address string) int64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("LogToDb:", err)
		}
	}()

	ret, err := MysqlDb.Exec("insert into chat_log (message,address)values(?,?)", msg, address)
	if err != nil {
		fmt.Println("LogToDb错误：", err)
		return -1
	}
	rows, _ := ret.RowsAffected()
	return rows
}
