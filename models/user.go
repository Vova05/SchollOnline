package models

type User struct {
	Id int `json:"-" `
	Username string `json:"username"`
	Password string `json:"password"`
	Token string `json:"token"`
}

type UpdateUser struct {
	Username string
	Password string
	Token string
}
