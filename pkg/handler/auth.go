package handler

import (
	"Online_school1/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)



func (h *Handler) signUp(c *gin.Context){

	var input models.User

	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id, err := h.service.Authorisation.CreateUser(input)
	if err != nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding: "required"`
	Password string `json:"password" binding: "required"`
}


func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	Username , err1 := c.GetPostForm("Username")
	Password, err2  := c.GetPostForm("Password")

	if err1 != true || err2 != true{
		newErrorResponse(c,http.StatusBadRequest, "auth failed")
		return
	}
	input.Username=Username
	input.Password=Password
	idUser, token, err := h.service.Authorisation.GenerateToken(input.Username, input.Password)
	if err != nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	log.Println(token)

	h.service.Authorisation.SaveToken(idUser,token)


	cookie := &http.Cookie{
		Name: "Token",
		Value: token,
	}
	http.SetCookie(c.Writer,cookie)
	c.Redirect(http.StatusMovedPermanently, "http://localhost:9091/school/profile")
}

func (h *Handler) signInGet(c *gin.Context){
	data := gin.H{
		"title": "Login",
	}
	c.HTML(http.StatusOK,"login1.html",data)
}


type login struct {
	Token string `json:"token"`
}