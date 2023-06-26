package models

import (
<<<<<<< HEAD
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/config"
	"github.com/jinzhu/gorm"
=======
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/config"
	"github.com/google/uuid"
>>>>>>> 3e408e5ce5de0504c7f5fe5c85eef5072b27d4bd
)



type Prayer struct {
	Base
	PrayerPoints       string `sql:"type:text;" json:"content"`
	ProgramID  uuid.UUID  `gorm:"type:uuid;column:prayer_foreign_key;not null;"`
	IsPublished      bool `json:"isPublished"`
	PublishedBy uuid.UUID `gorm:"type:uuid;column:user_foreign_key;not null;"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Prayer{})
}

func (b *Prayer) CreatePrayer() ( *Prayer, error){
	db.NewRecord(b)
	err:= db.Create(&b).Error
	return b, err
}


func (b *Prayer) UpdatePrayer() ( *Prayer, error){
	err:= db.Save(&b).Error
	return b, err
}
func GetAllPrayers() ([]Prayer, error) {
	var prayers []Prayer
	err:=db.Model(&Prayer{}).Preload("Testimonies").Preload("Prayers").Find(&prayers).Error
	if err !=nil{
		return nil, err
	}
	return prayers, nil
}

func GetPrayerById(Id int64) (*Prayer, error) {
	var getPrayer Prayer
	err := db.Where("ID=?", Id).Find(&getPrayer).Error
	return &getPrayer, err
}

func DeletePrayer(ID int64) (Prayer, error) {
	var prayer Prayer
	err :=db.Where("ID=?", ID).Delete(prayer).Error
	return prayer, err
}
