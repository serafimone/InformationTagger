package utils

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetInt64FieldFromRequest(r *http.Request, fieldName string) int64 {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars[fieldName], 10, 64)
	if err != nil {
		log.Fatalf(err.Error())
		return -1
	}
	return id
}
