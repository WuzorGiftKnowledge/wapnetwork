package routes

import (
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
}
