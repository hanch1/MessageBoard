package model

type Message struct {
	Id       int
	Uid      int
	UserName string `json:"userName"`
	Content  string `json:"content"`
	Time     string `json:"time"`
}
