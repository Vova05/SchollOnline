package repository

import (
	"Online_school1/models"
	"gorm.io/gorm"
	"log"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB{
	return &UserDB{db: db}
}

func (r *UserDB) CreateUser(user models.User)(error){
	row := r.db.Select("username","password").Create(&user)
	if err := row.Scan(&user).Error; err != nil {
		return err
	}
	return  nil
}

func (r *UserDB) UpdateUser(user models.User)(error){
	err := r.db.Table("users").Where("id = ?",user.Id).Updates(models.User{
		Username: user.Username,
		Password: user.Password,
		Token: user.Token,
	}).Error
	return err
}
func (r *UserDB) DeleteUser(idUser int)(error){
	var user models.User
	err := r.db.Table("users").Where("id = ?", idUser).Delete(user).Error
	return err
}
func (r *UserDB) GetAllUser()([]models.User){
	var usersTake []models.User
	err := r.db.Table("users").Find(&usersTake).Error
	if err != nil{
		log.Println(err.Error())
	}
	return usersTake
}
func (r *UserDB) GetUserById(idUser int)( models.User){
	var user models.User
	err := r.db.Table("users").Where("id = ?",idUser).Scan(&user).Error
	if err != nil {
		log.Println(err.Error())
	}
	return user
}

func (r *UserDB) GetByToken(token string)(int){
	var user models.User
	err := r.db.Table("users").Where("token = ?", token).First(&user).Error
	if err!=nil{
		log.Println(err.Error())
	}
	return user.Id
}

