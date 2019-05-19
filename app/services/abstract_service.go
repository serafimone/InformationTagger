package services

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/serafimone/InformationTagger/app/models"
	"github.com/serafimone/InformationTagger/app/utils"
	"github.com/serafimone/InformationTagger/config"
)

type AbstractService struct {
	Router       *mux.Router
	DBConnection *gorm.DB
}

func (service *AbstractService) Configure(config *config.AppConfig, router *mux.Router) {
	db := utils.GetConnection(config.DB)
	service.DBConnection = models.DBMigrate(db)
	service.Router = router
}

func (service *AbstractService) Run(host string) {
	log.Fatal(http.ListenAndServe(host, service.Router))
}
