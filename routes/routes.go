package routes

import (
	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-1/controllers"
	"github.com/gorilla/mux"
)

func MainRoutes() *mux.Router {

	r := mux.NewRouter().StrictSlash(false)
	//home handler to test wheather the Api is working or not
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	//Routes for the adding bullets
	r.HandleFunc("/api/magazine/addbullets", controllers.AddBullet).Methods("POST")

	//Routes for the Register magazine
	r.HandleFunc("/api/magazine/registermagazine", controllers.RegisterMagazine).Methods("POST")

	//Routes for the Register Gun
	r.HandleFunc("/api/magazine/registergun", controllers.RegisterGun).Methods("POST")

	// Routes for filling magazine
	r.HandleFunc("/api/magazine/fillmagazine", controllers.FillMagazine).Methods("POST")

	// Routes for Load magazine
	r.HandleFunc("/api/magazine/loadmagazine", controllers.LoadMagazine).Methods("POST")

	// Routes to Verify magazine
	r.HandleFunc("/api/magazine/verifymag", controllers.VerifyMag).Methods("POST")

	// Routes to Fire Gun
	r.HandleFunc("/api/magazine/firegun", controllers.FireGun).Methods("POST")

	return r

}
