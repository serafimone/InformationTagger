package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/serafimone/InformationTagger/app/core/authentication"
	"github.com/serafimone/InformationTagger/app/services"
)

func GetAuthAPIRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/token-auth", services.Login).Methods("POST")
	router.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(services.RefreshToken),
		)).Methods("GET")
	router.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(services.Logout),
		))
	return router
}
