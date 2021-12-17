package service

import (
	"Online_school1/models"
	"Online_school1/pkg/repository"
)

type LectureService struct {
	repo repository.Lecture
}

func NewLectureService(repo repository.Lecture) *LectureService{
	return &LectureService{repo: repo}
}

func (s *LectureService) Create(idGroup int,lec models.Lecture)(error){
	return s.repo.Create(idGroup,lec)
}
func (s *LectureService) GetAll()( []models.Lecture){
	return s.repo.GetAll()
}
func (s *LectureService) GetById(idLec int)( models.GetLecture){
	return s.repo.GetById(idLec)
}
func (s *LectureService) Update(idLec int, lecture models.Lecture)(error){
	return s.repo.Update(idLec,lecture)
}
func (s *LectureService) UpdateByGroup(idGroup int,lecId int)(error){
	return s.repo.UpdateByGroup(idGroup,lecId)
}
