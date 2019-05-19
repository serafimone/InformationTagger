package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/serafimone/InformationTagger/app/core/authentication"
	"github.com/serafimone/InformationTagger/app/core/parameters"
	"github.com/serafimone/InformationTagger/app/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	authBackend := authentication.InitJWTAuthenticationBackend()

	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.UUID)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, "")
		} else {
			response, _ := json.Marshal(parameters.TokenAuthentication{token})
			respondJSON(w, http.StatusOK, response)
		}
	}

	respondError(w, http.StatusUnauthorized, "")
}

func RefreshToken(w http.ResponseWriter, r *http.Request) []byte {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.UUID)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(w http.ResponseWriter, r *http.Request) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
	tokenString := r.Header.Get("Authorization")
	if err = authBackend.Logout(tokenString, tokenRequest); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
}
