package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context){


	cook , errC:= c.Request.Cookie("Token")

	var token string
	if errC == nil {
		token = cook.Value
	}else{
		newErrorResponse(c,http.StatusUnauthorized,"empty auth header")
		return
	}

	header := token
	if header == ""{
		newErrorResponse(c,http.StatusUnauthorized,"empty auth header")
		return
	}

	headerParts := strings.Split(header," ")
	log.Println(headerParts)
	log.Println(len(headerParts))


	userId, err := h.service.Authorisation.ParseToken(header)
	if err != nil {
		newErrorResponse(c,http.StatusUnauthorized,err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func (h *Handler)tokenParseToUserId ( nameCook string, c *gin.Context) (int,error){

	cook , errC:= c.Request.Cookie(nameCook)
	var token string
	if errC == nil {
		token = cook.Value
	}else{
		newErrorResponse(c,http.StatusUnauthorized,"empty auth header")
		return 0, errC
	}

	header := token
	if header == ""{
		newErrorResponse(c,http.StatusUnauthorized,"empty auth header")
		return 0, errC
	}

	headerParts := strings.Split(header," ")
	log.Println(headerParts)
	log.Println(len(headerParts))


	userId, err := h.service.Authorisation.ParseToken(header)
	if err != nil {
		newErrorResponse(c,http.StatusUnauthorized,err.Error())
		return 0,err
	}

	return userId,nil
}

func getUserId(c *gin.Context)(int ,error){
	id, ok := c.Get(userCtx)
	if  !ok  {
		newErrorResponse(c, http.StatusInternalServerError,"user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok{
		newErrorResponse(c,http.StatusInternalServerError,"user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
