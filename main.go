package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/serafimone/InformationTagger/app/core/container"
	"github.com/serafimone/InformationTagger/app/routes"
	"github.com/serafimone/InformationTagger/config"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS", "UPDATE"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	config := config.GetConfig()
	router := mux.NewRouter()
	router = routes.GetAuthAPIRoutes(router)
	router = routes.GetRESTAPIRoutes(router)
	container := container.Init(config.DB, router)
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(container.Router)
	n.Run(":3000")
}
