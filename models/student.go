package models

type Student struct {
	Id int `json:"id"`
	Stu_id int `json:"stu_id"`
	Surname string `json:"surname"`
	GroupId int `json:"group_id"`
}

type UpdateStudent struct {
	Stu_id int `json:"stu_id"`
	Surname string `json:"surname"`
	GroupId int `json:"group_id"`
}
