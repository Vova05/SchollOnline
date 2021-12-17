package handler

import (
	"Online_school1/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createStudent(c *gin.Context){

	//userId, err := h.tokenParseToUserId("Token",c)
	var input models.User
	var inputSt models.Student
	input.Username, _ = c.GetPostForm("Username")
	input.Password, _ =c.GetPostForm("Password")
	inputSt.Surname, _ = c.GetPostForm("Surname")
	GroupId, _ := c.GetPostForm("GroupId")
	inputSt.GroupId, _ = strconv.Atoi(GroupId)

	h.service.Student.Create(input,inputSt)


	c.Redirect(http.StatusMovedPermanently, "http://localhost:9091/school/students")
}

type GetInfoAboutStudents struct {
	Surname string
	GroupId int
	Login string
}

func (h *Handler) GetAllStudentTeacher (c *gin.Context){
	userId, _ := h.tokenParseToUserId("Token",c)
	userData := h.service.User.GetUserById(userId)
	_,students := h.service.Student.GetAll()
	groups := h.service.Group.GetAll()
	var studentsTeacher []models.Student
	var groupsTeacher []models.Group
	var usersTeacher []models.User
	for index, _:= range groups{
		if groups[index].IdTeacher == userId{
			groupsTeacher=append(groupsTeacher,groups[index])
		}
	}

	for index,_:=range students{
		for index2,_:=range groupsTeacher{
			if groupsTeacher[index2].Id==students[index].GroupId{
				studentsTeacher=append(studentsTeacher,students[index])
				usersTeacher=append(usersTeacher,h.service.GetUserById(students[index].Stu_id))
			}
		}
	}

	var infoFront []GetInfoAboutStudents
	for index,_:=range usersTeacher{
		var tmpInfo GetInfoAboutStudents
		tmpInfo.Surname = studentsTeacher[index].Surname
		tmpInfo.GroupId = studentsTeacher[index].GroupId
		tmpInfo.Login = usersTeacher[index].Username
		infoFront = append(infoFront,tmpInfo)
	}

	data := gin.H{
		"title": "Students",
		"NameUser": userData.Username,
		"Students": infoFront,
		"Role": "Teacher",
	}
	c.HTML(http.StatusOK,"students.html",data)
}

func (h *Handler) GetSort (c *gin.Context){
	userId, _ := h.tokenParseToUserId("Token",c)
	userData := h.service.User.GetUserById(userId)
	_,students := h.service.Student.SortByName()
	groups := h.service.Group.GetAll()
	var studentsTeacher []models.Student
	var groupsTeacher []models.Group
	var usersTeacher []models.User
	for index, _:= range groups{
		if groups[index].IdTeacher == userId{
			groupsTeacher=append(groupsTeacher,groups[index])
		}
	}

	for index,_:=range students{
		for index2,_:=range groupsTeacher{
			if groupsTeacher[index2].Id==students[index].GroupId{
				studentsTeacher=append(studentsTeacher,students[index])
				usersTeacher=append(usersTeacher,h.service.GetUserById(students[index].Stu_id))
			}
		}
	}

	var infoFront []GetInfoAboutStudents
	for index,_:=range usersTeacher{
		var tmpInfo GetInfoAboutStudents
		tmpInfo.Surname = studentsTeacher[index].Surname
		tmpInfo.GroupId = studentsTeacher[index].GroupId
		tmpInfo.Login = usersTeacher[index].Username
		infoFront = append(infoFront,tmpInfo)
	}

	data := gin.H{
		"title": "Students",
		"NameUser": userData.Username,
		"Students": infoFront,
		"Role": "Teacher",
	}
	c.HTML(http.StatusOK,"students.html",data)
}
