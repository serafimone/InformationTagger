package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/serafimone/InformationTagger/app/models"
)

func AddRecord(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	record, err := models.CreateRecord(db, r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, record)
}

func GetAllRecords(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	records, err := models.GetAllRecords(db, r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, records)
}

func GetRecord(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	record, err := models.GetRecord(db, r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, record)
}

func DeleteRecord(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	err := models.DeleteRecord(db, r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, "Success")
}

func UpdateRecordContent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	record, err := models.UpdateRecordContent(db, r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, record)
}
