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

//App is the instance of application
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

//Initialize is app initialization
func (app *App) Initialize(config *config.AppConfig) {
	db := utils.GetConnection(config.DB)
	app.DB = models.DBMigrate(db)
	app.Router = mux.NewRouter()
}

// Run the app on it's router
func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}

func (app *App) setRouters() {
	router := app.Router
	router.HandleFunc("/"+utils.APIVersion+"/documents", app.getAllDocuments).Methods("GET")
}

func (app *App) getAllDocuments(w http.ResponseWriter, r *http.Request) {
	handlers.GetAllDocuments(app.DB, w, r)
}
