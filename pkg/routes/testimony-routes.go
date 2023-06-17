package routes

import (
	"github.com/WuzorGiftKnowledge/testimonyapp/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterTestimonyStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/testimony/", controllers.CreateTestimony).Methods("POST")
	router.HandleFunc("/testimony/", controllers.GetTestimony).Methods("GET")
	router.HandleFunc("/testimony/{testimonyId}", controllers.GetTestimonyById).Methods("GET")
	router.HandleFunc("/testimony/{testimonyId}", controllers.UpdateTestimony).Methods("PUT")
	router.HandleFunc("/testimony/{testimonyId}", controllers.DeleteTestimony).Methods("DELETE")
}
