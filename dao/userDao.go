package dao

import (
	db "app1/database"
	"app1/model"
)

func GetUserByNameAndPassword(username string, password string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password from user where username = ? and password = ?"
	//执行
	row := db.SqlDB.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password)
	return user, nil
}

func GetUserByName(username string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password from user where username = ?"
	//执行
	row := db.SqlDB.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password)
	return user, nil
}

//SaveUser 向数据库中插入用户信息
func SaveUser(username string, password string) error {
	//写sql语句
	sqlStr := "insert into user(username,password) values(?,?)"
	//执行
	_, err := db.SqlDB.Exec(sqlStr, username, password)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUser() ([]*model.User, error) {
	sqlStr := "select id,username from user"
	rows, err := db.SqlDB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for rows.Next() {
		var user = &model.User{}
		rows.Scan(&user.Id, &user.Username)
		users = append(users, user)
	}
	return users, nil
}

func GetUsernameByUid(uid int) (string, error) {
	sqlStr := "select * from user where id = ?"
	row := db.SqlDB.QueryRow(sqlStr, uid)
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password)
	return user.Username, nil
}
