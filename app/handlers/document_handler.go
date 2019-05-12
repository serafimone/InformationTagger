package handlers

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/serafimone/InformationTagger/app/models"
)

//GetAllDocuments tries to get all documents from database
func GetAllDocuments(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	documents, err := models.GetAllDocuments(db, w, r)
	if documents == nil || err != nil {
		fmt.Println(err.Error())
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, documents)
}

//GetDocument tries to get document from database
func GetDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	document, err := models.GetDocument(db, w, r)
	if document == nil || err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, document)
}

//CreateDocument creates document and try insert it to database
func CreateDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	document, err := models.CreateDocument(db, w, r)
	if document == nil || err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, document)

}

//DeleteDocument tries to delete document from database
func DeleteDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	err := models.DeleteDocument(db, w, r)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}
