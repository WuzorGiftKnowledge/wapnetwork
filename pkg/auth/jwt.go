package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"strings"
	"time"

	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/models"
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

var secret = utils.GetEnvVariable("jwtSecretKey")

func loginHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var message Message
	err := json.NewDecoder(request.Body).Decode(&message)
	if err != nil {
		return
	}
	err = json.NewEncoder(writer).Encode(message)
	if err != nil {
		return
	}
}

func GenerateJWTToken(username string, id uuid.UUID) (accessToken string, rt string, err error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(15 * time.Minute).Unix()
	claims["authorized"] = true
	claims["current_username"] = username
	claims["current_user_id"] = id

	if secret == "" {
		return "", "", fmt.Errorf("unable to create token")
	}
	accessToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", "", fmt.Errorf(err.Error())
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["current_user_id"] = id
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	rt, err = refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", "", fmt.Errorf(err.Error())
	}

	return accessToken, rt, nil
}

func AuthMiddleware(endpointHandler http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authHeader := strings.Split(request.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]

			token, err := verifyToken(jwtToken)
			if err != nil {

				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte(err.Error()))
				return

			}
			if token.Valid {
				claims, ok := token.Claims.(jwt.MapClaims)
				if ok {
					ctx := context.WithValue(request.Context(), "props", claims)
					endpointHandler.ServeHTTP(writer, request.WithContext(ctx))
				}
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("You're Unauthorized due to invalid token"))
				if err != nil {
					return
				}
			}
		}

	})
}

func RefreshToken(refreshToken string) (accessToken string, rt string, err error) {

	token, err := verifyToken(refreshToken)
	if err != nil {

		return "", "", err
	}

	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)

		if ok && claims["current_user_id"] != nil {
			userid := claims["current_user_id"].(float64)

			user, db := models.GetUserById(int64(userid))
			if db.Error != nil {
				err = fmt.Errorf("user not found")
				return
			}

			return GenerateJWTToken(user.Email, uuid.UUID(user.ID))
		}
	}

	return "", "", fmt.Errorf("token is not invalid token")

}

func verifyToken(jwtToken string) (token *jwt.Token, err error) {

	token, err = jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			err = fmt.Errorf("Malformed token")
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			err = fmt.Errorf("Token is expired")
		} else {
			err = fmt.Errorf("failed with error: %s", err.Error())
		}
	}
	return
}
