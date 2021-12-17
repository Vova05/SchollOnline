package service

import (
	"Online_school1/models"
	"Online_school1/pkg/repository"
)

type GroupService struct {
	repo repository.Group
}

func NewGroupService(repo repository.Group) *GroupService{
	return &GroupService{repo: repo}
}

func (s *GroupService) Create(group models.Group)(error){
	return s.repo.Create(group)
}
func (s *GroupService) GetAll()( []models.Group){
	return s.repo.GetAll()
}
func (s *GroupService) GetById(idGroup int)( models.Group){
	return s.repo.GetById(idGroup)
}
func (s *GroupService) Update(idGroup int, group models.Group)(error){
	return s.repo.Update(idGroup,group)
}

func (s *GroupService) Delete(idGroup int)(error){
	return s.repo.Delete(idGroup)
}
