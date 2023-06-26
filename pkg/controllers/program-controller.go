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

var NewProgram models.Program

func GetProgram(w http.ResponseWriter, r *http.Request) {
	newPrograms, err := models.GetAllPrograms()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error calling GetProgram:%s", err.Error())))
		return
	}
	res, err := json.Marshal(newPrograms)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while marshalling "))
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProgramById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	programId := vars["programId"]
	ID, err := uuid.Parse(programId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while parsing id: %s", err.Error())))
		return
	}
	programDetails, err := models.GetProgramById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error calling GetProgramById:%s", err.Error())))
		return
	}
	res, _ := json.Marshal(programDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateProgram(w http.ResponseWriter, r *http.Request) {
	CreateProgram := &models.Program{}
	CurrentUser, ok := r.Context().Value("values").(models.CurrentUser)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while retrieving user details from context"))
		return
	}

	CreateProgram.CreatedBy = uuid.UUID(CurrentUser.Id)
	utils.ParseBody(r, CreateProgram)
	b, err := CreateProgram.CreateProgram()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error calling CreateProgram:%s", err.Error())))
		return
	}
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProgram(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	programId := vars["programId"]
	ID, err := uuid.Parse(programId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while parsing id: %s", err.Error())))
		return
	}
	program, err := models.GetProgramById(ID)
	if err != nil || program == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error calling DeleteProgram:%s", err.Error())))
		return
	}
	err = program.DeleteProgram(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error calling DeleteProgram:%s", err.Error())))
		return
	}
	res, _ := json.Marshal(program)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProgram(w http.ResponseWriter, r *http.Request) {
	var updateProgram = &models.Program{}
	utils.ParseBody(r, updateProgram)
	vars := mux.Vars(r)
	programId := vars["programId"]
	ID, err := uuid.Parse(programId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while parsing id: %s", err.Error())))
		return
	}
	programDetails, err := models.GetProgramById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error calling GetProgram:%s", err.Error())))
		return
	}
	if updateProgram.Name != "" {
		programDetails.Name = updateProgram.Name
	}
	if updateProgram.Description != "" {
		programDetails.Description = updateProgram.Description
	}
	if updateProgram.Date.IsZero() {
		programDetails.Date = updateProgram.Date
	}

	programDetails.UpdateProgram()
	res, err := json.Marshal(programDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while marshalling "))
		return 
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
