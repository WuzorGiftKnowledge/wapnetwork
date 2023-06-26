package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/auth"
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/models"
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/utils"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := &models.LoginRequest{}

	utils.ParseBody(r, loginRequest)
	user, err := models.GetUserByUserName(loginRequest.UserName)
	if err != nil || user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("incorrect username"))
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid credentials"))
		return
	}

	accessToken, refreshToken, err := auth.GenerateJWTToken(user.Username, uuid.UUID(user.ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error generating token " + err.Error()))
		return
	}
	loginResponse := &models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	res, _ := json.Marshal(loginResponse)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshRequest := &models.RefreshTokenRequest{}

	utils.ParseBody(r, refreshRequest)

	accessToken, refreshToken, err := auth.RefreshToken(refreshRequest.ResfreshToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error refreshing token " + err.Error()))
		return
	}
	loginResponse := &models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	res, _ := json.Marshal(loginResponse)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
