package service

import (
	"Online_school1/models"
	"Online_school1/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService{
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.User)(error){
	user.Password= generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user models.User)(error){
	user.Password = generatePasswordHash(user.Password)
	return s.repo.UpdateUser(user)
}
func (s *UserService) DeleteUser(idUser int)(error){
	return s.repo.DeleteUser(idUser)
}
func (s *UserService) GetAllUser()( []models.User){
	return s.repo.GetAllUser()
}
func (s *UserService) GetUserById(idUser int)( models.User){
	return s.repo.GetUserById(idUser)
}

func (s *UserService) GetByToken(token string)(int){
	return s.repo.GetByToken(token)
}
