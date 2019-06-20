package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"

	"github.com/jinzhu/gorm"
	"github.com/serafimone/InformationTagger/app/models"
	"github.com/serafimone/InformationTagger/app/requests"
	"github.com/unidoc/unioffice/document"
)

// GetAllDocuments tries to get all documents from database
func GetAllDocuments(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	documents, err := models.GetAllDocuments(db, r)
	if documents == nil || err != nil {
		fmt.Println(err.Error())
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, documents)
}

// GetDocument tries to get document from database
func GetDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	document, err := models.GetDocument(db, r)
	if document == nil || err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, document)
}

// CreateDocument creates document and try insert it to database
func CreateDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	document, err := models.CreateDocument(db, w, r)
	if document == nil || err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, document)

}

// DeleteDocument tries to delete document from database
func DeleteDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	err := models.DeleteDocument(db, w, r)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, "Success")
}

func UpdateDocumentTitle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	document, err := models.UpdateDocumentTitle(db, w, r)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, document)
}

func FormDocument(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	body := r.Body
	request := requests.FormDocumentRequest{}
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&request); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	documentFromDatabase, err := models.GetDocumentFromDatabase(request.DocumentID, db)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	dox := document.New()
	for _, element := range documentFromDatabase.Records {
		paragraph := dox.AddParagraph()
		run := paragraph.AddRun()
		run.AddText(element.Content)
		run.Properties().SetSize(measurement.Distance(request.FontSize))
		run.Properties().SetFontFamily(request.Font)
		paragraph.Properties().Spacing().SetLineSpacing(measurement.Distance(request.FontSize*request.Interval), wml.ST_LineSpacingRuleAuto)
	}
	dox.SaveToFile("D:/file.docx")
	respondJSON(w, http.StatusOK, "Success")
}
