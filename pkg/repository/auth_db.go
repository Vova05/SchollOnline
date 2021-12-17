package repository

import (
	"Online_school1/models"
	"gorm.io/gorm"
)

type AuthDB struct {
	db *gorm.DB
}

func NewAuthDB(db *gorm.DB) *AuthDB{
	return &AuthDB{db: db}
}

func  (r *AuthDB)SaveToken(userId int,token string)(error){
	var user models.User
	r.db.Table("users").Where("id = ?",userId).First(&user)
	user.Id = userId
	err := r.db.Table("users").Where("id = ?",userId).Update("token", token).Error
	if  err != nil {
		return err
	}
	return nil
}

func (r *AuthDB) TakeToken(userId int)(string,error){
	var user models.User
	err := r.db.Table("users").Where("id = ?",userId).First(&user).Error
	token := user.Token
	return token, err
}
func (r *AuthDB) CreateUser(user models.User)(int,error){
	var id int
	row := r.db.Select("Username","Password").Create(&user)
	if err := row.Scan(&user).Error; err != nil {
		return 0,err
	}
	id = user.Id
	return id, nil
}

func (r *AuthDB) GetUser(username,password string)(int,models.User,error){
	var user models.User
	err := r.db.Where("username = ? AND password = ?",username,password).First(&user).Error
	return user.Id,user, err
}
