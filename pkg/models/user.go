package models

import (
	"github.com/WuzorGiftKnowledge/bookapp/pkg/config"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"" json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}
func GetUserByEmail(email string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("email=?", email).Find(&getUser)
	return &getUser, db
}
func DeleteUser(ID int64) User {
	var User User
	db.Where("ID=?", ID).Delete(User)
	return User
}
