package service

import (
	"Online_school1/models"
	"Online_school1/pkg/repository"
)

type StudentService struct {
	repo repository.Student
}

func NewStudentService(repo repository.Student) *StudentService{
	return &StudentService{repo: repo}
}

func (s *StudentService) Create(user models.User, student models.Student)(error){
	return s.repo.Create(user,student)
}
func (s *StudentService) GetAll()( []models.User,  []models.Student){
	return s.repo.GetAll()
}
func (s *StudentService) GetById(idStudent int)( models.User,  models.Student){
	return s.repo.GetById(idStudent)
}
func (s *StudentService) Update(idUser int,student models.Student)(error){
	return s.repo.Update(idUser,student)
}

func (s *StudentService) SortByName()( []models.User,  []models.Student)  {
	return s.repo.SortByName()
}
