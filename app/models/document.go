package models

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Document type represents documents
type Document struct {
	gorm.Model
	Title   string   `gorm:"unique" json:"title"`
	Records []Record `gorm:"ForeignKey:DocumentID" json:"records"`
}

//GetAllDocuments try to get all documents from database
func GetAllDocuments(db *gorm.DB, w http.ResponseWriter, r *http.Request) (*[]Document, error) {
	documents := []Document{}
	err := db.Find(&documents)
	return &documents, err.Error
}

//GetDocument try to get document from database
func GetDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) (*Document, error) {
	vars := mux.Vars(r)
	title := vars["title"]
	return getDocumentFromDatabase(title, db, w)
}

//CreateDocument create document and insert it into database
func CreateDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) (*Document, error) {
	document := Document{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&document)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	err = db.Save(&document).Error
	return &document, nil
}

//DeleteDocument try to delete document from database
func DeleteDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	title := vars["title"]
	document, err := getDocumentFromDatabase(title, db, w)
	if document == nil || err != nil {
		return err
	}
	db.Delete(&document)
	return nil
}

func getDocumentFromDatabase(title string, db *gorm.DB, w http.ResponseWriter) (*Document, error) {
	document := Document{}
	if err := db.First(&document, Document{Title: title}).Error; err != nil {
		return nil, err
	}
	return &document, nil
}
