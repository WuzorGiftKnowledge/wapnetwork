package routes

import (
	"github.com/WuzorGiftKnowledge/bookapp/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/signup", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/refresh_token", controllers.RefreshToken).Methods("POST")
	
}
