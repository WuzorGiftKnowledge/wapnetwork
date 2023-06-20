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

var NewPrayer models.Prayer

func GetPrayer(w http.ResponseWriter, r *http.Request) {
	newPrayers, err := models.GetAllPrayers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling GetPrayer:%s",err.Error())))
			return
	}
	res, err:= json.Marshal(newPrayers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while marshalling "))
			return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPrayerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prayerId := vars["prayerId"]
	ID, err := strconv.ParseInt(prayerId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while parsing id"))
			return
	}
	prayerDetails, err := models.GetPrayerById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling GetPrayerById:%s",err.Error())))
			return
	}
	res, _ := json.Marshal(prayerDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatePrayer(w http.ResponseWriter, r *http.Request) {
	CreatePrayer := &models.Prayer{}
	utils.ParseBody(r, CreatePrayer)
	b, err := CreatePrayer.CreatePrayer()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling CreatePrayer:%s",err.Error())))
			return
	}
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePrayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prayerId := vars["prayerId"]
	ID, err := strconv.ParseInt(prayerId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while parsing id"))
			return
	}
	prayer,err := models.DeletePrayer(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling DeletePrayer:%s",err.Error())))
			return
	}
	res, _ := json.Marshal(prayer)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdatePrayer(w http.ResponseWriter, r *http.Request) {
	var updatePrayer = &models.Prayer{}
	utils.ParseBody(r, updatePrayer)
	vars := mux.Vars(r)
	prayerId := vars["prayerId"]
	ID, err := strconv.ParseInt(prayerId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while parsing"))
			return
	}
	prayerDetails, err := models.GetPrayerById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling GetPrayer:%s",err.Error())))
			return
	}
	if !prayerDetails.IsPublished {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Can not update a published record."))
		return
	}
	if updatePrayer.PrayerPoints != "" {
		prayerDetails.PrayerPoints = updatePrayer.PrayerPoints
	}
	if updatePrayer.ProgramID.String() != "" {
		prayerDetails.ProgramID = updatePrayer.ProgramID
	}
	prayerDetails.UpdatePrayer()
	res, _ := json.Marshal(prayerDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func PublishPrayer(w http.ResponseWriter, r *http.Request) {
	var updatePrayer = &models.Prayer{}
	utils.ParseBody(r, updatePrayer)
	vars := mux.Vars(r)
	prayerId := vars["prayerId"]
	ID, err := strconv.ParseInt(prayerId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while parsing"))
			return
		
	}
	prayerDetails, err := models.GetPrayerById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error calling GetPrayerById:%s",err.Error())))
			return
	}
	if prayerDetails.IsPublished != false {
		prayerDetails.IsPublished =true
	}else{
		prayerDetails.IsPublished =false
	}

	prayerDetails.UpdatePrayer()
	res, err := json.Marshal(prayerDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while marshalling "))
			return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
