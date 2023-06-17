package models

import (
	"github.com/WuzorGiftKnowledge/prayerapp/pkg/config"
	"github.com/jinzhu/gorm"
)



type Prayer struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Prayer{})
}

func (b *Prayer) CreatePrayer() *Prayer {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllPrayers() []Prayer {
	var Prayers []Prayer
	db.Find(&Prayers)
	return Prayers
}

func GetPrayerById(Id int64) (*Prayer, *gorm.DB) {
	var getPrayer Prayer
	db := db.Where("ID=?", Id).Find(&getPrayer)
	return &getPrayer, db
}

func DeletePrayer(ID int64) Prayer {
	var prayer Prayer
	db.Where("ID=?", ID).Delete(prayer)
	return prayer
}
