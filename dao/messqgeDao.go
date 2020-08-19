package dao

import (
	db "app1/database"
	"app1/model"
)

// 返回留言message
func GetMessages(pageIndex int, pageSize int) ([]*model.Message, error) {
	limit := (pageIndex - 1) * pageSize

	//写sql语句
	sqlStr := "select m.id, u.username, m.content, m.time from message m, user u where m.uid = u.id limit ?,? "
	//执行
	rows, err := db.SqlDB.Query(sqlStr, limit, pageSize)
	if err != nil {
		return nil, err
	}
	var msgs []*model.Message
	for rows.Next() {
		msg := &model.Message{}
		rows.Scan(&msg.Id, &msg.UserName, &msg.Content, &msg.Time)
		msgs = append(msgs, msg)
	}
	return msgs, nil
}

func AddMessage(msg *model.Message) error {
	slqStr := "insert into message(uid,content,time) values(?,?,?)"
	_, err := db.SqlDB.Exec(slqStr, msg.Uid, msg.Content, msg.Time)
	if err != nil {
		return err
	}
	return nil
}
