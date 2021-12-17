package repository

import (
	"Online_school1/models"
	"gorm.io/gorm"
)

type Authorisation interface {
	CreateUser(user models.User)(int,error)
	GetUser(username,password string)(int,models.User,error)
	SaveToken(userId int,token string)(error)
	TakeToken(userId int)(string,error)
}

type User interface {
	CreateUser(user models.User)(error)
	UpdateUser(user models.User)(error)
	DeleteUser(idUser int)(error)
	GetAllUser()( []models.User)
	GetUserById(idUser int)( models.User)
	GetByToken(token string)(int)
}

type Student interface {
	Create(user models.User, student models.Student)(error)
	GetAll()( []models.User,  []models.Student)
	GetById(idStudent int)( models.User,  models.Student)
	Update(idUser int,student models.Student)(error)
	SortByName( )( []models.User,  []models.Student)
}

type Group interface {
	Create(group models.Group)(error)
	GetAll()( []models.Group)
	GetById(idGroup int)( models.Group)
	Update(idGroup int, group models.Group)(error)
	Delete(idGroup int)(error)

}

type Lecture interface {
	Create(idGroup int,lec models.Lecture)(error)
	GetAll()( []models.Lecture)
	GetById(idLec int)( models.GetLecture)
	Update(idLec int, lecture models.Lecture)(error)
	UpdateByGroup(idGroup int,lecId int)(error)
}

type Repository struct {
	Authorisation
	User
	Student
	Group
	Lecture
}

func NewRepository(db *gorm.DB) *Repository{
	return &Repository{
		Authorisation: NewAuthDB(db),
		User: NewUserDB(db),
		Student: NewStudentDB(db),
		Group: NewGroupDB(db),
		Lecture: NewLectureDB(db),
	}
}