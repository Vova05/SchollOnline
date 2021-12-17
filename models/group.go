package models

type Group struct {
	Id int `json:"id"`
	IdTeacher int `json:"id_teacher"`
}

type UpdateGroup struct {
	IdTeacher int `json:"id_teacher"`
}
