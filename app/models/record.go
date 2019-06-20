package models

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/serafimone/InformationTagger/app/utils"
)

// Record type represents records
type Record struct {
	gorm.Model
	Content    string `json:"Content"`
	DocumentID uint   `json:"DocumentID"`
}

func CreateRecord(db *gorm.DB, r *http.Request) (*Record, error) {
	document, err := GetDocument(db, r)
	if err != nil {
		return nil, err
	}
	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer r.Body.Close()
	content := string(requestData)
	record := Record{Content: content, DocumentID: document.ID}
	context := db.Model(&document).Related(&[]Record{}).Save(&record)
	return &record, context.Error
}

func GetAllRecords(db *gorm.DB, r *http.Request) (*[]Record, error) {
	document, err := GetDocument(db, r)
	if err != nil {
		return nil, err
	}
	records := []Record{}
	context := db.Model(&document).Related(&records)
	return &records, context.Error
}

func GetRecord(db *gorm.DB, r *http.Request) (*Record, error) {
	document, err := GetDocument(db, r)
	if err != nil {
		return nil, err
	}
	id := utils.GetInt64FieldFromRequest(r, "record_id")
	record := Record{}
	context := db.Model(&document).Related(&[]Record{}).Where([]int64{id}).First(&record)
	return &record, context.Error
}

func DeleteRecord(db *gorm.DB, r *http.Request) error {
	document, err := GetDocument(db, r)
	if err != nil {
		return err
	}
	id := utils.GetInt64FieldFromRequest(r, "record_id")
	context := db.Model(&document).Related(&[]Record{}).Where([]int64{id}).Delete(&Record{})
	return context.Error
}

func UpdateRecordContent(db *gorm.DB, r *http.Request) (*Record, error) {
	document, err := GetDocument(db, r)
	if err != nil {
		return nil, err
	}
	id := utils.GetInt64FieldFromRequest(r, "record_id")
	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer r.Body.Close()
	content := string(requestData)
	record := Record{}
	context := db.Model(&document).Related(&[]Record{}).Where([]int64{id}).First(&record)
	if context.Error != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	record.Content = content
	context = db.Save(&record)
	return &record, context.Error
}

func getDocumentRecords(db *gorm.DB, document *Document) {
	if err := db.Model(document).Related(&document.Records); err == nil {
		panic("Error, while setting document records!")
	}
}
