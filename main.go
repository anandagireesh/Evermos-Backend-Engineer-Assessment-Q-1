package main

import (
	"net/http"

	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-1/database"
	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-1/routes"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func main() {

	database.DbConnection()
	database.Db.Ping()

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.UseHandler(routes.MainRoutes())

	server := &http.Server{
		Addr:    "0.0.0.0:8006",
		Handler: n,
	}

	log.Info("Server Running")

	server.ListenAndServe()

}
