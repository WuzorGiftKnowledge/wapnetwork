package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/models"
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/utils"
	"github.com/gorilla/mux"
)

var NewTestimony models.Testimony

func GetTestimony(w http.ResponseWriter, r *http.Request) {
	newTestimonys, err := models.GetAllTestimonys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling GetTestimony:%s",err.Error())))
			return
	}
	res, err:= json.Marshal(newTestimonys)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while marshalling "))
			return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTestimonyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testimonyId := vars["testimonyId"]
	ID, err := strconv.ParseInt(testimonyId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while parsing id"))
			return
	}
	testimonyDetails, err := models.GetTestimonyById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling GetTestimonyById:%s",err.Error())))
			return
	}
	res, _ := json.Marshal(testimonyDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTestimony(w http.ResponseWriter, r *http.Request) {
	CreateTestimony := &models.Testimony{}
	utils.ParseBody(r, CreateTestimony)
	b, err := CreateTestimony.CreateTestimony()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling CreateTestimony:%s",err.Error())))
			return
	}
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTestimony(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testimonyId := vars["testimonyId"]
	ID, err := strconv.ParseInt(testimonyId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while parsing id"))
			return
	}
	testimony,err := models.DeleteTestimony(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling DeleteTestimony:%s",err.Error())))
			return
	}
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
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while parsing"))
			return
	}
	testimonyDetails, err := models.GetTestimonyById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling GetTestimony:%s",err.Error())))
			return
	}
	if !testimonyDetails.IsPublished {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Can not update a published record."))
		return
	}
	if updateTestimony.Content != "" {
		testimonyDetails.Content = updateTestimony.Content
	}
	if updateTestimony.ProgramID.String() != "" {
		testimonyDetails.ProgramID = updateTestimony.ProgramID
	}
	testimonyDetails.UpdateTestimony()
	res, _ := json.Marshal(testimonyDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func PublishTestimony(w http.ResponseWriter, r *http.Request) {
	var updateTestimony = &models.Testimony{}
	utils.ParseBody(r, updateTestimony)
	vars := mux.Vars(r)
	testimonyId := vars["testimonyId"]
	ID, err := strconv.ParseInt(testimonyId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while parsing"))
			return
		
	}
	testimonyDetails, err := models.GetTestimonyById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling GetTestimonyById:%s",err.Error())))
			return
	}
	if testimonyDetails.IsPublished != false {
		testimonyDetails.IsPublished =true
	}else{
		testimonyDetails.IsPublished =false
	}

	testimonyDetails.UpdateTestimony()
	res, err := json.Marshal(testimonyDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while marshalling "))
			return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
