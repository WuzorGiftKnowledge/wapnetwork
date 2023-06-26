package models

import (
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/config"
	"github.com/google/uuid"
)



type Testimony struct {
	Base
	Content        string `sql:"type:text;" json:"content"`
	ProgramID  uuid.UUID  `gorm:"not null" json:"programID"`
	IsPublished      bool `json:"isPublished"`
	PublishedBy uuid.UUID `json:"publishedBy"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Testimony{})
}

func (b *Testimony) CreateTestimony() ( *Testimony, error){
	db.NewRecord(b)
	err:= db.Create(&b).Error
	return b, err
}


func (b *Testimony) UpdateTestimony() ( *Testimony, error){
	err:= db.Save(&b).Error
	return b, err
}
func GetAllTestimonys() ([]Testimony, error) {
	var testimonys []Testimony
	err:=db.Model(&Testimony{}).Find(&testimonys).Error
	if err !=nil{
		return nil, err
	}
	return testimonys, nil
}

func GetTestimonyById(Id uuid.UUID) (*Testimony, error) {
	var getTestimony Testimony
	err := db.Where("ID=?", Id).Find(&getTestimony).Error
	return &getTestimony, err
}

func DeleteTestimony(ID uuid.UUID) (Testimony, error) {
	var testimony Testimony
	err :=db.Where("ID=?", ID).Delete(testimony).Error
	return testimony, err
}
