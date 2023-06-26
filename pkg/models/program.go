package models

import (
	"time"
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/config"
	"github.com/google/uuid"
	
)



type Program struct {
	Base
	Name       string `gorm:"not null" json:"name"`
	Description     string `gorm:"" json:"description"`
	Date     time.Time `gorm:"not null" json:"date"`
	Testimonies []Testimony `json:"testimonies"`
	Prayers []Prayer `json:"prayers"`
	CreatedBy uuid.UUID `gorm:"not null"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Program{})
}

func (b *Program) CreateProgram() ( *Program, error){
	db.NewRecord(b)
	err:= db.Create(&b).Error
	return b, err
}


func (b *Program) UpdateProgram() ( *Program, error){
	err:= db.Save(&b).Error
	return b, err
}
func GetAllPrograms() ([]Program, error) {
	var programs []Program
	err:=db.Model(&Program{}).Preload("Testimonies").Preload("Prayers").Find(&programs).Error
	if err !=nil{
		return nil, err
	}
	return programs, nil
}

func GetProgramById(Id uuid.UUID) (*Program, error) {
	var getProgram Program
	err := db.Where("ID=?", Id).Find(&getProgram).Error
	return &getProgram, err
}

func (b *Program) DeleteProgram(ID uuid.UUID) error {
	var program Program
	err := db.Where("ID=?", ID).Delete(program).Error
	return err
}
