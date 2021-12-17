package models

type GroupHomework struct {
	Id int `json:"id"`
	HomId int `json:"hom_id"`
	GroupId int `json:"group_id"`
}

type Homework struct {
	Id int `json:"id"`
	Task string `json:"task"`
	Example string `json:"example"`
}

type UpdateHomework struct {
	Task string `json:"task"`
	Example string `json:"example"`
}
