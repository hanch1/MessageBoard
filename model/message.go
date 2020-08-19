package model

type Message struct {
	Id       int
	UserName string `json:"userName"`
	Content  string `json:"content"`
	Time     string `json:"time"`
}
