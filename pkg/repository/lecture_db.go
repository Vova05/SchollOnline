package repository

import (
	"Online_school1/models"
	"gorm.io/gorm"
	"log"
)

type LectureDB struct {
	db *gorm.DB
}

func NewLectureDB(db *gorm.DB) *LectureDB{
	return &LectureDB{db: db}
}

func (r *LectureDB) Create(idGroup int,lec models.Lecture)(error){
	err := r.db.Table("lecture").Create(&lec).Error
	err1 := r.db.Table("lecture").Where("link",lec.Link).First(&lec).Error
	var lec_group models.GroupLecture
	lec_group.GroupId=idGroup
	lec_group.LecId=lec.Id
	err2 := r.db.Table("group_lecture").Create(&lec_group).Error
	if  err != nil || err1!=nil || err2!=nil{
		log.Println(err.Error())
		log.Println(err1.Error())
		log.Println(err2.Error())
		return err
	}
	return  nil
}
func (r *LectureDB) GetAll()( []models.Lecture){
	var lectures []models.Lecture
	err := r.db.Table("lecture").Find(&lectures).Error
	if err != nil{
		log.Println(err.Error())
	}
	return lectures
}
func (r *LectureDB) GetById(idLec int)( models.GetLecture){
	var lec models.Lecture
	err := r.db.Table("lecture").Where("id = ?",idLec).Scan(&lec).Error
	if err != nil {
		log.Println(err.Error())
	}
	var group_lec []models.GroupLecture
	err2 := r.db.Table("group_lecture").Where("lec_id",idLec).Find(&group_lec).Error
	if err2 != nil {
		log.Println(err.Error())
	}
	var id_groups []int
	for index,_:=range group_lec{
		id_groups=append(id_groups,group_lec[index].GroupId)
	}
	var lecGroup models.GetLecture
	lecGroup.GroupId = id_groups
	lecGroup.Time = lec.Time
	lecGroup.Link = lec.Link
	lecGroup.Id = lec.Id
	return lecGroup
}
func (r *LectureDB) Update(idLec int, lecture models.Lecture)(error){
	err := r.db.Table("lecture").Where("id = ?",idLec).Updates(models.Lecture{
		Time: lecture.Time,
		Link: lecture.Link,
	}).Error
	return err
}

func (r *LectureDB) UpdateByGroup(idGroup int,lecId int)(error){
	var group_lec models.GroupLecture
	group_lec.LecId=lecId
	group_lec.GroupId=idGroup
	row := r.db.Table("group_lecture").Create(&group_lec)
	if err := row.Scan(&group_lec).Error; err != nil {
		return err
	}
	return  nil
}
