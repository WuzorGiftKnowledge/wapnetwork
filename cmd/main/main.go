package main

import (
	"log"
	"net/http"

	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/auth"
	"github.com/WuzorGiftKnowledge/wapnetwork/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")

	r := mux.NewRouter()

	authrouter := r.PathPrefix("/api/auth").Subrouter()
	routes.RegisterAuthRoutes(authrouter)

	apirouter := r.PathPrefix("/api").Subrouter()
	routes.RegisterProgramRoutes(apirouter)
	routes.RegisterPrayerStoreRoutes(apirouter)
	routes.RegisterTestimonyStoreRoutes(apirouter)
	routes.RegisterUserRoutes(apirouter)
	
	apirouter.Use(auth.AuthMiddleware)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
