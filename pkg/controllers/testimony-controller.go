package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/models"
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var NewTestimony models.Testimony

func GetTestimony(w http.ResponseWriter, r *http.Request) {
	newTestimonys, err := models.GetAllTestimonys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error calling GetTestimony:%s", err.Error())))
		return
	}
	res, err := json.Marshal(newTestimonys)
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
	ID, err := uuid.Parse(testimonyId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while parsing id: %s", err.Error())))
		return
	}
	testimonyDetails, err := models.GetTestimonyById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error calling GetTestimonyById:%s", err.Error())))
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
		w.Write([]byte(fmt.Sprintf("error calling CreateTestimony:%s", err.Error())))
		return
	}
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTestimony(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testimonyId := vars["testimonyId"]
	ID, err := uuid.Parse(testimonyId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while parsing id: %s", err.Error())))
		return
	}
	testimony, err := models.DeleteTestimony(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error calling DeleteTestimony:%s", err.Error())))
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
	ID, err := uuid.Parse(testimonyId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while parsing id: %s", err.Error())))
		return
	}
	testimonyDetails, err := models.GetTestimonyById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error calling GetTestimony:%s", err.Error())))
		return
	}
	if testimonyDetails.IsPublished {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Can not update a published record. Unplishe the record first"))
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
	ID, err := uuid.Parse(testimonyId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while parsing id: %s", err.Error())))
		return
	}
	testimonyDetails, err := models.GetTestimonyById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error calling GetTestimonyById:%s", err.Error())))
		return
	}
	CurrentUser, ok := r.Context().Value("values").(models.CurrentUser)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while retrieving user details from context"))
		return
	}
	if testimonyDetails.IsPublished == false {
		testimonyDetails.IsPublished = true
		testimonyDetails.PublishedBy = uuid.UUID(CurrentUser.Id)
	} else {
		testimonyDetails.IsPublished = false
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
