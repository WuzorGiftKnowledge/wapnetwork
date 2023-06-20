package routes

import (
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterProgramRoutes = func(router *mux.Router) {
	router.HandleFunc("/program/", controllers.CreateProgram).Methods("POST")
	router.HandleFunc("/program/", controllers.GetProgram).Methods("GET")
	router.HandleFunc("/program/{programId}", controllers.GetProgramById).Methods("GET")
	router.HandleFunc("/program/{programId}", controllers.UpdateProgram).Methods("PUT")
	router.HandleFunc("/program/{programId}", controllers.DeleteProgram).Methods("DELETE")
}
