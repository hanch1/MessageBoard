package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

// 初始化数据库连接
func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/app1?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	}
	// 最大连接数
	SqlDB.SetMaxOpenConns(20)
	// 最大空闲连接数
	SqlDB.SetMaxIdleConns(20)
}
