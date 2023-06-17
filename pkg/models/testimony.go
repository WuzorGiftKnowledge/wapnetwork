package models

import (
	"github.com/WuzorGiftKnowledge/testimonyapp/pkg/config"
	"github.com/jinzhu/gorm"
)



type Testimony struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Testimony{})
}

func (b *Testimony) CreateTestimony() *Testimony {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllTestimonys() []Testimony {
	var testimonys []Testimony
	db.Find(&testimonys)
	return testimonys
}

func GetTestimonyById(Id int64) (*Testimony, *gorm.DB) {
	var getTestimony Testimony
	db := db.Where("ID=?", Id).Find(&getTestimony)
	return &getTestimony, db
}

func DeleteTestimony(ID int64) Testimony {
	var testimony Testimony
	db.Where("ID=?", ID).Delete(testimony)
	return testimony
}
