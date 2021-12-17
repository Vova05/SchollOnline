package repository

import (
	"Online_school1/models"
	"gorm.io/gorm"
	"log"
)

type StudentDB struct {
	db *gorm.DB
}

func NewStudentDB(db *gorm.DB) *StudentDB{
	return &StudentDB{db: db}
}

func (r *StudentDB) Create(user models.User, student models.Student)(error){
	var userCreate models.User
	err := r.db.Select("username","password").Create(&user).Error
	err2 := r.db.Table("users").Where("username = ?", user.Username).First(&userCreate).Error
	student.Stu_id=userCreate.Id
	err3 := r.db.Table("students").Create(&student).Error
	if err != nil || err2 != nil || err3!=nil {
		log.Println(err)
		log.Println(err2)
		log.Println(err3)
		return err
	}
	return  nil
}
func (r *StudentDB) GetAll()( []models.User,  []models.Student){
	var usersTake []models.User
	err := r.db.Table("users").Find(&usersTake).Error
	if err != nil{
		log.Println(err.Error())
	}
	var studentTake []models.Student
	//for _, _ = range usersTake {
	//	var student models.Student
	//	err2 := r.db.Table("students").Scan(&student).Error
	//	if err2!= nil{
	//		log.Println(err2.Error())
	//	}
	//	studentTake = append(studentTake,student)
	//}
	 r.db.Table("students").Find(&studentTake)
	return usersTake,studentTake
}
func (r *StudentDB) GetById(idStudent int)( models.User,  models.Student){
	var user models.User
	var student models.Student
	err := r.db.Table("students").Where("id = ?",idStudent).Scan(&student).Error
	idUser:=student.Stu_id
	err2 := r.db.Table("users").Where("id = ?",idUser).Scan(&user).Error
	if err != nil || err2 != nil{
		log.Println(err.Error())
		log.Println(err2.Error())
	}
	return user,student
}
func (r *StudentDB) Update(idUser int,student models.Student)(error){
	err := r.db.Table("student").Where("stu_id = ?",idUser).Updates(models.Student{
		Stu_id: student.Stu_id,
		Surname: student.Surname,
		GroupId: student.GroupId,
	}).Error
	return err
}

func (r *StudentDB) SortByName( )( []models.User,  []models.Student){
	var studentTake []models.Student
	r.db.Table("students").Order("surname").Find(&studentTake)
	var tmp models.User
	var usersTake []models.User
	for index,_:= range studentTake{
		r.db.Table("users").Where("id = ?",studentTake[index].Stu_id).Scan(&tmp)
		usersTake = append(usersTake,tmp)
	}
	return usersTake,studentTake
}
