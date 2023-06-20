package models

import (
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/config"
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
)



type Program struct {
	Base
	Name       string `json:"name"`
	Description     string `json:"description"`
	Date     time.Time `json:"date"`
	Testimonies []Testimony `json:"testimonies"`
	Prayers []Prayer `json:"prayers"`
	CreatedBy uuid.UUID `gorm:"type:uuid;column:user_foreign_key;not null;"`
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

func GetProgramById(Id int64) (*Program, error) {
	var getProgram Program
	err := db.Where("ID=?", Id).Find(&getProgram).Error
	return &getProgram, err
}

func DeleteProgram(ID int64) (Program, error) {
	var program Program
	err :=db.Where("ID=?", ID).Delete(program).Error
	return program, err
}
