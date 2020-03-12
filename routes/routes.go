package routes

import (
	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-1/controllers"
	"github.com/gorilla/mux"
)

func MainRoutes() *mux.Router {

	r := mux.NewRouter().StrictSlash(false)
	//home handler to test wheather the Api is working or not
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	//Routes for the magazine
	r.HandleFunc("/api/magazine/addbullets", controllers.AddBullet).Methods("POST")

	return r

}
