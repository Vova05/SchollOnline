package repository

import (
	"Online_school1/models"
	"gorm.io/gorm"
	"log"
	"time"
)

type ApplicationDB struct {
	db *gorm.DB
}

type order struct {
	IdOrder       int
	IdAdmin       int
	IdClient      int
	IdWorkshop    int
	IdOrderStatus int
	DateStart     time.Time
	DateEnd       time.Time
}

type client struct {
	IdClient int
	FCs      string
	Passport string
	Number   string
}

type admin struct {
	IdAdmin  int
	FCs      string
	Passport string
	Number   string
}

type orderStatus struct {
	FCs      string
	Passport string
	Number   string
}

type workshop struct {
	IdWorkshop  int
	IdEquipment int
	Description string
	Square      int
}

type price struct {
	IdPrice    int
	IdWorkshop int
	DailyPrice int
	Sale       int
}

func NewApplicationDB(db *gorm.DB) *ApplicationDB {
	return &ApplicationDB{db: db}
}

func (r *ApplicationDB) Create(idGroup int, lec models.Application) error {

	return nil
}
func (r *ApplicationDB) GetAll() []models.Application {
	var applications []models.Application

	var orderM []order
	var workshopM []workshop
	var clientM []client
	var adminM []admin

	err := r.db.Table("order").Find(&orderM).Error

	for i := range orderM {
		err1 := r.db.Table("workshop").Where("id_workshop", orderM[i].IdWorkshop).Find(&workshopM)
		err2 := r.db.Table("client").Where("id_client", orderM[i].IdClient).Find(&clientM)
		err3 := r.db.Table("admin").Where("id_admin", orderM[i].IdAdmin).Find(&adminM)
		if err1 == nil && err2 == nil && err3 == nil {
			var apNp models.Application
			apNp.NameClient = clientM[i].FCs
			apNp.NameAdmin = adminM[i].FCs
			apNp.DateEnd = orderM[i].DateEnd.String()
			apNp.DateStart = orderM[i].DateStart.String()
			apNp.IdRequest = orderM[i].IdOrder
			apNp.IdWorkshop = workshopM[i].IdWorkshop
			applications = append(applications, apNp)
		}

	}
	if err != nil {
		log.Println(err.Error())
	}
	return applications
}
func (r *ApplicationDB) GetById(idLec int) models.Application {
	var ap models.Application
	return ap
}
func (r *ApplicationDB) Update(idLec int, lecture models.Application) error {
	return nil
}
