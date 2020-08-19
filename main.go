package main

import (
	db "app1/database"
	."app1/router"
)

func main() {
	// 方法执行后关闭数据库连接
	defer db.SqlDB.Close()
	router := InitRouter()
	router.Run()
}

