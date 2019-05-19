package services

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/serafimone/InformationTagger/app/core/container"
	"github.com/serafimone/InformationTagger/app/handlers"
)

// DOCUMENTS API BLOCK

func GetAllDocuments(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.GetAllDocuments(container.GetDBConnection(), w, r)
}

func GetDocument(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.GetDocument(container.GetDBConnection(), w, r)
}

func CreateDocument(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.CreateDocument(container.GetDBConnection(), w, r)
}

func DeleteDocument(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.DeleteDocument(container.GetDBConnection(), w, r)
}

func UpdateDocumentTitle(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.UpdateDocumentTitle(container.GetDBConnection(), w, r)
}

// RECORDS API BLOCK

func GetAllDocumentRecords(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.GetAllRecords(container.GetDBConnection(), w, r)
}

func AddRecordToDocument(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.AddRecord(container.GetDBConnection(), w, r)
}

func GetDocumentRecord(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.GetRecord(container.GetDBConnection(), w, r)
}

func UpdateDocumentRecord(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.UpdateRecordContent(container.GetDBConnection(), w, r)
}

func DeleteDocumentRecord(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	handlers.DeleteRecord(container.GetDBConnection(), w, r)
}
