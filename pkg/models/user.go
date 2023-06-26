package models

import (
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/config"
	
)

type User struct {
	Base
	FirstName string `gorm:"not null" json:"firstname"`
	LastName  string `gorm:"not null" json:"lastname"`
	Email     string `gorm:"unique;not null;type:varchar(100)" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Username  string `gorm:"unique;not null;type:varchar(100)" json:"username"`
}


func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser()  (*User, error){
	db.NewRecord(b)
	err:=db.Create(&b).Error
	return b, err
}
func (b *User) UpdateUser()  error{
	err:= db.Save(&b).Error
	return  err
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, error) {
	var getUser User
	err := db.Where("ID=?", Id).Find(&getUser).Error
	return &getUser, err
}
func GetUserByEmail(email string) (*User, error) {
	var getUser User
	err := db.Where("email=?", email).Find(&getUser).Error
	return &getUser, err
}
func GetUserByUserName(username string) (*User, error) {
	var getUser User
	err := db.Where("username=?", username).Find(&getUser).Error
	return &getUser, err
}
func DeleteUser(ID int64) (User, error) {
	var User User
	err:= db.Where("ID=?", ID).Delete(User).Error
	return User, err
}

