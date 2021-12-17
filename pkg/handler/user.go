package handler

import (
	"Online_school1/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
)

func (h *Handler) createUser(c *gin.Context){

	var input models.User
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	err := h.service.User.CreateUser(input)

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"status": "ok",
	})
}//JSON

func (h *Handler) updateUser(c *gin.Context){
//изменить
	var input models.User
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}


	err := h.service.User.UpdateUser(input)

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"status": "ok",
	})
}//JSON

type GetGroup struct {
	Id int
	Students []models.Student
}
func (h *Handler) getProfileTeacher (c *gin.Context){

	userId, _ := h.tokenParseToUserId("Token",c)
	userData := h.service.User.GetUserById(userId)
	Groups := h.service.Group.GetAll()
	var GroupsUser []models.Group
	for index, _:= range Groups{
		if Groups[index].IdTeacher==userId{
			GroupsUser=append(GroupsUser,Groups[index])
		}
	}
	var getGroups []GetGroup
	_ ,students := h.service.Student.GetAll()
	for index, _:=range GroupsUser{
		var tmpGetGroup GetGroup
		tmpGetGroup.Id=GroupsUser[index].Id
		for index2,_:= range students{
			if students[index2].GroupId==GroupsUser[index].Id{
				tmpGetGroup.Students = append(tmpGetGroup.Students,students[index2])
			}
		}
		getGroups = append(getGroups,tmpGetGroup)
	}
	data := gin.H{
		"title": "Profile",
		"NameUser": userData.Username,
		"Groups": getGroups,
		"Role": "Teacher",
	}
	c.HTML(http.StatusOK,"profile.html",data)//поставить свой html
}

func (h *Handler) getGroupSortById (c *gin.Context){
	userId, _ := h.tokenParseToUserId("Token",c)
	userData := h.service.User.GetUserById(userId)
	Groups := h.service.Group.GetAll()
	var GroupsUser []models.Group
	for index, _:= range Groups{
		if Groups[index].IdTeacher==userId{
			GroupsUser=append(GroupsUser,Groups[index])
		}
	}

	sort.SliceStable(GroupsUser, func(i, j int) bool {
		return GroupsUser[i].Id>GroupsUser[i].Id
	})
	data := gin.H{
		"title": "Profile",
		"NameUser": userData.Username,
		"Groups": GroupsUser,
	}
	c.HTML(http.StatusOK,"profile.html",data)
}

func (h *Handler) createGroup (c *gin.Context){
	userId, _ := h.tokenParseToUserId("Token",c)
	 var group models.Group
	group.IdTeacher = userId
	h.service.Group.Create(group)
	c.Redirect(http.StatusMovedPermanently, "http://localhost:9091/school/profile")
}

func (h *Handler) deleteGroup (c *gin.Context){
	Id,_ := c.GetPostForm("id_delete")
	IdGroup, _ := strconv.Atoi(Id)
	h.service.Group.Delete(IdGroup)
	c.Redirect(http.StatusMovedPermanently, "http://localhost:9091/school/profile")
}



