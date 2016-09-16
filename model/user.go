package model

type User struct{
	Id int64 `json:"id"`
	Username string  `json:"username"`
	Mail string  `json:"mail"`
	Password string `json:"password"` 
}

type MessageSend struct {
	Message   string   `json:"message"`
	Consumers []string `json:"consumers"`
}