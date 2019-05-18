package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/serafimone/InformationTagger/app/handlers"
	"github.com/serafimone/InformationTagger/app/models"
	"github.com/serafimone/InformationTagger/app/utils"
	"github.com/serafimone/InformationTagger/config"
)

// App is the instance of application
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize is app initialization
func (app *App) Initialize(config *config.AppConfig) {
	db := utils.GetConnection(config.DB)
	app.DB = models.DBMigrate(db)
	app.Router = mux.NewRouter()
	app.setRouters()
}

// Run the app on it's router
func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}

func (app *App) setRouters() {
	router := app.Router
	router.HandleFunc(utils.APIVersion+"documents", app.getAllDocuments).Methods("GET")
	router.HandleFunc(utils.APIVersion+"documents/{document_id}", app.getDocument).Methods("GET")
	router.HandleFunc(utils.APIVersion+"documents/document", app.createDocument).Methods("POST")
	router.HandleFunc(utils.APIVersion+"documents/{document_id}", app.updateDocumentTitle).Methods("POST")
	router.HandleFunc(utils.APIVersion+"documents/{document_id}", app.deleteDocument).Methods("DELETE")
	router.HandleFunc(utils.APIVersion+"documents/{document_id}/records", app.getAllDocumentRecords).Methods("GET")
	router.HandleFunc(utils.APIVersion+"documents/{document_id}/records/{record_id}", app.getDocumentRecord).Methods("GET")
	router.HandleFunc(utils.APIVersion+"documents/{document_id}/record", app.addRecordToDocument).Methods("POST")
	router.HandleFunc(utils.APIVersion+"documents/{document_id}/records/{record_id}", app.updateDocumentRecord).Methods("POST")
	router.HandleFunc(utils.APIVersion+"documents/{document_id}/records/{record_id}", app.deleteDocumentRecord).Methods("DELETE")
}

// DOCUMENTS API BLOCK

func (app *App) getAllDocuments(w http.ResponseWriter, r *http.Request) {
	handlers.GetAllDocuments(app.DB, w, r)
}

func (app *App) getDocument(w http.ResponseWriter, r *http.Request) {
	handlers.GetDocument(app.DB, w, r)
}

func (app *App) createDocument(w http.ResponseWriter, r *http.Request) {
	handlers.CreateDocument(app.DB, w, r)
}

func (app *App) deleteDocument(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteDocument(app.DB, w, r)
}

func (app *App) updateDocumentTitle(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateDocumentTitle(app.DB, w, r)
}

// RECORDS API BLOCK

func (app *App) getAllDocumentRecords(w http.ResponseWriter, r *http.Request) {
	handlers.GetAllRecords(app.DB, w, r)
}

func (app *App) addRecordToDocument(w http.ResponseWriter, r *http.Request) {
	handlers.AddRecord(app.DB, w, r)
}

func (app *App) getDocumentRecord(w http.ResponseWriter, r *http.Request) {
	handlers.GetRecord(app.DB, w, r)
}

func (app *App) updateDocumentRecord(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateRecordContent(app.DB, w, r)
}

func (app *App) deleteDocumentRecord(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteRecord(app.DB, w, r)
}
