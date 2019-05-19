package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/serafimone/InformationTagger/app/core/container"
	"github.com/serafimone/InformationTagger/app/routes"
	"github.com/serafimone/InformationTagger/config"
)

func main() {
	config := config.GetConfig()
	router := mux.NewRouter()
	router = routes.GetAuthAPIRoutes(router)
	router = routes.GetRESTAPIRoutes(router)
	container := container.Init(config.DB, router)
	n := negroni.Classic()
	n.UseHandler(container.Router)
	http.ListenAndServe(":3000", n)

}
