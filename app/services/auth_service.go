package services

import (
	"net/http"

	"github.com/serafimone/InformationTagger/app/handlers"
)

func Login(w http.ResponseWriter, r *http.Request) {
	handlers.Login(w, r)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.RefreshToken(w, r)
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.Logout(w, r)
}
