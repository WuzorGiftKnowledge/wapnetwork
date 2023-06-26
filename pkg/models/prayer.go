package models

import (
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/config"
	"github.com/google/uuid"
	
)

type Prayer struct {
	Base
	PrayerPoints string    `gorm:"type:text;" json:"prayerPoints"`
	ProgramID    uuid.UUID `gorm:"not null" json:"programID"`
	IsPublished  bool      `json:"isPublished"`
	PublishedBy  uuid.UUID `json:"publishedBy"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Prayer{})
}

func (b *Prayer) CreatePrayer() (*Prayer, error) {
	db.NewRecord(b)
	err := db.Create(&b).Error
	return b, err
}

func (b *Prayer) UpdatePrayer() (*Prayer, error) {
	err := db.Save(&b).Error
	return b, err
}
func GetAllPrayers() ([]Prayer, error) {
	var prayers []Prayer
	err := db.Model(&Prayer{}).Find(&prayers).Error
	if err != nil {
		return nil, err
	}
	return prayers, nil
}

func GetPrayerById(Id uuid.UUID) (*Prayer, error) {
	var getPrayer Prayer
	err := db.Where("ID=?", Id).Find(&getPrayer).Error
	return &getPrayer, err
}

func DeletePrayer(ID uuid.UUID) (Prayer, error) {
	var prayer Prayer
	err := db.Where("ID=?", ID).Delete(prayer).Error
	return prayer, err
}
