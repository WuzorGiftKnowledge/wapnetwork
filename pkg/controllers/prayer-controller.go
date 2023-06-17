package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/WuzorGiftKnowledge/prayerapp/pkg/models"
	"github.com/WuzorGiftKnowledge/prayerapp/pkg/utils"
	"github.com/gorilla/mux"
)

var NewPrayer models.Prayer

func GetPrayer(w http.ResponseWriter, r *http.Request) {
	newPrayers := models.GetAllPrayers()
	res, _ := json.Marshal(newPrayers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPrayerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prayerId := vars["prayerId"]
	ID, err := strconv.ParseInt(prayerId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	prayerDetails, _ := models.GetPrayerById(ID)
	res, _ := json.Marshal(prayerDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatePrayer(w http.ResponseWriter, r *http.Request) {
	CreatePrayer := &models.Prayer{}
	utils.ParseBody(r, CreatePrayer)
	b := CreatePrayer.CreatePrayer()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePrayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prayerId := vars["prayerId"]
	ID, err := strconv.ParseInt(prayerId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	prayer := models.DeletePrayer(ID)
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
		fmt.Println("error while parsing")
	}
	prayerDetails, db := models.GetPrayerById(ID)
	if updatePrayer.Name != "" {
		prayerDetails.Name = updatePrayer.Name
	}
	if updatePrayer.Author != "" {
		prayerDetails.Author = updatePrayer.Author
	}
	if updatePrayer.Publication != "" {
		prayerDetails.Publication = updatePrayer.Publication
	}
	db.Save(&prayerDetails)
	res, _ := json.Marshal(prayerDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
