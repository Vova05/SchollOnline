package models

import "time"

type GroupLecture struct {
	Id int `json:"id"`
	LecId int `json:"lec_id"`
	GroupId int `json:"group_id"`
}

type Lecture struct {
	Id int `json:"id"`
	Time time.Time `json:"time"`
	Link string `json:"link"`
}

type GetLecture struct {
	Id int `json:"id"`
	Time time.Time `json:"time"`
	Link string `json:"link"`
	GroupId []int `json:"group_id"`
}

type UpdateLecture struct {
	Time time.Time `json:"time"`
	Link string `json:"link"`
}
