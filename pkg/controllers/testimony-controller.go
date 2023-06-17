package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/WuzorGiftKnowledge/testimonyapp/pkg/models"
	"github.com/WuzorGiftKnowledge/testimonyapp/pkg/utils"
	"github.com/gorilla/mux"
)

var NewTestimony models.Testimony

func GetTestimony(w http.ResponseWriter, r *http.Request) {
	newTestimonys := models.GetAllTestimonys()
	res, _ := json.Marshal(newTestimonys)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTestimonyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testimonyId := vars["testimonyId"]
	ID, err := strconv.ParseInt(testimonyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	testimonyDetails, _ := models.GetTestimonyById(ID)
	res, _ := json.Marshal(testimonyDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTestimony(w http.ResponseWriter, r *http.Request) {
	CreateTestimony := &models.Testimony{}
	utils.ParseBody(r, CreateTestimony)
	b := CreateTestimony.CreateTestimony()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTestimony(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testimonyId := vars["testimonyId"]
	ID, err := strconv.ParseInt(testimonyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	testimony := models.DeleteTestimony(ID)
	res, _ := json.Marshal(testimony)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTestimony(w http.ResponseWriter, r *http.Request) {
	var updateTestimony = &models.Testimony{}
	utils.ParseBody(r, updateTestimony)
	vars := mux.Vars(r)
	testimonyId := vars["testimonyId"]
	ID, err := strconv.ParseInt(testimonyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	testimonyDetails, db := models.GetTestimonyById(ID)
	if updateTestimony.Name != "" {
		testimonyDetails.Name = updateTestimony.Name
	}
	if updateTestimony.Author != "" {
		testimonyDetails.Author = updateTestimony.Author
	}
	if updateTestimony.Publication != "" {
		testimonyDetails.Publication = updateTestimony.Publication
	}
	db.Save(&testimonyDetails)
	res, _ := json.Marshal(testimonyDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
