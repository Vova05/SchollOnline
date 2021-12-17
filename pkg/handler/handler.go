package handler

import (
	"Online_school1/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.StaticFS("/css", http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\Online_school1\\templates\\css"))
	router.StaticFS("/fonts", http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\Online_school1\\templates\\css\\fonts"))
	router.StaticFS("/img", http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\Online_school1\\templates\\img"))
	router.StaticFS("/js", http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\Online_school1\\templates\\js"))
	router.StaticFS("/less", http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\Online_school1\\templates\\less"))

	router.LoadHTMLGlob("C:\\Users\\VovaGlh\\GolandProjects\\Online_school1\\templates/*.html")


	router.POST("/auths",h.signIn)
	router.POST("/new",h.signUp)
	router.GET("/auth",h.signInGet)
	router.PUT("/profileUpdate",h.updateUser)
	group := router.Group("/school",h.userIdentity)
	{
		group.StaticFS("/css", http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\Online_school1\\templates\\css"))
		group.StaticFS("/fonts", http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\Online_school1\\templates\\css\\fonts"))
		group.StaticFS("/img", http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\Online_school1\\templates\\img"))
		group.StaticFS("/js", http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\Online_school1\\templates\\js"))
		group.StaticFS("/less", http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\Online_school1\\templates\\less"))

		group.GET("/profile",h.getProfileTeacher)
		group.POST("/add_group",h.createGroup)
		group.POST("/delete",h.deleteGroup)

		group.GET("/sort_name",h.GetSort)
		group.POST("/create_st",h.createStudent)

		//group.GET("/group",)
		//group.POST("/group",)

		group.GET("/students",h.GetAllStudentTeacher)
		group.POST("/student",)

	}

	return router
}
