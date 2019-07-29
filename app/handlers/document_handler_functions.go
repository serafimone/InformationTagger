package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"

	"github.com/jinzhu/gorm"
	"github.com/serafimone/InformationTagger/app/models"
	"github.com/serafimone/InformationTagger/app/requests"
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
	models.GetDocumentRecords(db, documentFromDatabase)
	dox := document.New()
	for _, element := range documentFromDatabase.Records {
		paragraph := dox.AddParagraph()
		run := paragraph.AddRun()
		run.AddText(element.Content)
		run.Properties().SetSize(measurement.Distance(request.FontSize))
		run.Properties().SetFontFamily(request.Font)
		paragraph.Properties().Spacing().SetLineSpacing(measurement.Distance(request.FontSize*request.Interval), wml.ST_LineSpacingRuleAtLeast)
	}
	fileName := documentFromDatabase.Title + ".docx"
	dox.SaveToFile(fileName)
	sendFile(fileName, w)
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func sendFile(path string, w http.ResponseWriter) {
	extraParams := map[string]string{
		"title":       "My Document",
		"author":      "Matt Aimonetti",
		"description": "A document with all the Go programming language secrets",
	}
	request, err := newfileUploadRequest("https://api.anonymousfiles.io/", extraParams, "file", path)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		respondError(w, http.StatusInternalServerError, err.Error())
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)
		fmt.Println(body)
		respondJSON(w, http.StatusOK, body.String())
	}
}
