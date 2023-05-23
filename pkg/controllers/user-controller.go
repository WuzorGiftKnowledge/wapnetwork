package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"log"
	"github.com/WuzorGiftKnowledge/bookapp/pkg/models"
	"github.com/WuzorGiftKnowledge/bookapp/pkg/utils"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()

	for _, u := range users {
		sanitizeUser(&u)
	}
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId := vars["UserId"]
	ID, err := strconv.ParseInt(UserId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, _ := models.GetUserById(ID)
	sanitizeUser(userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	createUser := &models.User{}
	utils.ParseBody(r, createUser)
	ok := isUserExist(createUser.Email)
	

	if ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("User already exists with the given email"))
		return
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(createUser.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("please consider another password"))
		return
	}
	createUser.Password = string(encryptedPassword)

	createUser.Username = createUser.Email

	u := createUser.CreateUser()
	sanitizeUser(u)
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId := vars["UserId"]
	ID, err := strconv.ParseInt(UserId, 0, 0)
	if err != nil {
		log.Println("error while parsing")
	}
	user := models.DeleteUser(ID)
	sanitizeUser(&user)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	UserId := vars["UserId"]
	ID, err := strconv.ParseInt(UserId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	UserDetails, db := models.GetUserById(ID)
	if updateUser.Email != "" || updateUser.Email != UserDetails.Email {
		ok := isUserExist(updateUser.Email)
		if ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("User already exists with the given email"))
			return
		}
		UserDetails.Email = updateUser.Email
	}
	if updateUser.FirstName != "" {
		UserDetails.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		UserDetails.LastName = updateUser.LastName
	}
	db.Save(&UserDetails)
	sanitizeUser(UserDetails)
	res, _ := json.Marshal(UserDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func isUserExist(emailorusername string) bool {

	_, err := models.GetUserByEmail(emailorusername)

	if err.Error == nil {
		return true
	}
	_, err = models.GetUserByUserName(emailorusername)
	if err.Error == nil {
		return true

	}
	return false
}

func sanitizeUser(u *models.User) {
	if u != nil {
		u.Password = ""
	}

}
