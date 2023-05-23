package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/WuzorGiftKnowledge/bookapp/pkg/models"
	"github.com/WuzorGiftKnowledge/bookapp/pkg/utils"
	"github.com/gorilla/mux"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUsers()
	res, _ := json.Marshal(newUsers)
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
	UserDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(UserDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId := vars["UserId"]
	ID, err := strconv.ParseInt(UserId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	User := models.DeleteUser(ID)
	res, _ := json.Marshal(User)
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
	if updateUser.Email != "" {
		UserDetails.Email = updateUser.Email
	}
	if updateUser.FirstName != "" {
		UserDetails.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		UserDetails.LastName = updateUser.LastName
	}
	db.Save(&UserDetails)
	res, _ := json.Marshal(UserDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
