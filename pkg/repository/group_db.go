package repository

import (
	"Online_school1/models"
	"gorm.io/gorm"
	"log"
)

type GroupDB struct {
	db *gorm.DB
}

func NewGroupDB(db *gorm.DB) *GroupDB{
	return &GroupDB{db: db}
}



func (r *GroupDB) Create(group models.Group)(error){
	err := r.db.Table("group").Create(&group).Error
	if  err != nil {
		return err
	}
	return  nil
}
func (r *GroupDB) GetAll()([]models.Group){
	var groups []models.Group
	err := r.db.Table("group").Find(&groups).Error
	if err != nil{
		log.Println(err.Error())
	}
	return groups
}
func (r *GroupDB) GetById(idGroup int)( models.Group){
	var group models.Group
	err := r.db.Table("group").Where("id = ?",idGroup).Scan(&group).Error
	if err != nil {
		log.Println(err.Error())
	}
	return group
}
func (r *GroupDB) Update(idGroup int, group models.Group)(error){
	err := r.db.Table("group").Where("id = ?",idGroup).Updates(models.Group{
		IdTeacher: group.IdTeacher,
	}).Error
	return err
}

func (r *GroupDB) Delete(idGroup int)(error){
	var g models.Group
	err := r.db.Table("group").Where("id = ?", idGroup).Delete(&g)
	return err.Error
}
