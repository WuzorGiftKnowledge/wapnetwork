package routes

import (
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterPrayerStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/prayer/", controllers.CreatePrayer).Methods("POST")
	router.HandleFunc("/prayer/", controllers.GetPrayer).Methods("GET")
	router.HandleFunc("/prayer/{prayerId}", controllers.GetPrayerById).Methods("GET")
	router.HandleFunc("/prayer/{prayerId}", controllers.UpdatePrayer).Methods("PUT")
	router.HandleFunc("/prayer/publish/{prayerId}", controllers.PublishPrayer).Methods("PUT")
	router.HandleFunc("/prayer/{prayerId}", controllers.DeletePrayer).Methods("DELETE")
}
