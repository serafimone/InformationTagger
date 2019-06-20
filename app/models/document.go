package models

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/serafimone/InformationTagger/app/utils"
)

//Document type represents documents
type Document struct {
	gorm.Model
	Title   string   `gorm:"unique" json:"Title"`
	Records []Record `gorm:"ForeignKey:DocumentID" json:"Records"`
}

//GetAllDocuments try to get all documents from database
func GetAllDocuments(db *gorm.DB, r *http.Request) (*[]Document, error) {
	documents := []Document{}
	context := db.Find(&documents)
	for index := range documents {
		getDocumentRecords(db, &documents[index])
	}
	return &documents, context.Error
}

//GetDocument try to get document from database
func GetDocument(db *gorm.DB, r *http.Request) (*Document, error) {
	document := Document{}
	documentID := utils.GetInt64FieldFromRequest(r, "document_id")
	context := db.Where([]int64{documentID}).First(&document)
	getDocumentRecords(db, &document)
	return &document, context.Error
}

//CreateDocument creates document and insert it into database
func CreateDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) (*Document, error) {
	document := Document{}
	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer r.Body.Close()
	title := string(requestData)
	document.Title = title
	err = db.Save(&document).Error
	return &document, nil
}

func UpdateDocumentTitle(db *gorm.DB, w http.ResponseWriter, r *http.Request) (*Document, error) {
	id := utils.GetInt64FieldFromRequest(r, "document_id")
	document, err := GetDocumentFromDatabase(id, db)
	if err != nil {
		return nil, err
	}
	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer r.Body.Close()
	title := string(requestData)
	document.Title = title
	context := db.Save(&document)
	return document, context.Error
}

//DeleteDocument try to delete document from database
func DeleteDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) error {
	id := utils.GetInt64FieldFromRequest(r, "document_id")
	document, err := GetDocumentFromDatabase(id, db)
	if document == nil || err != nil {
		return err
	}
	db.Delete(&document)
	return nil
}

func GetDocumentFromDatabase(id int64, db *gorm.DB) (*Document, error) {
	document := Document{}
	if err := db.First(&document, id).Error; err != nil {
		return nil, err
	}
	return &document, nil
}
