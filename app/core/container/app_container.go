package container

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/serafimone/InformationTagger/app/models"
	"github.com/serafimone/InformationTagger/app/utils"
	"github.com/serafimone/InformationTagger/config"
)

type AppContainer struct {
	DB     *gorm.DB
	Router *mux.Router
}

var container *AppContainer = nil

func Init(config *config.DBConfig, router *mux.Router) *AppContainer {
	if container == nil {
		db := utils.GetConnection(config)
		models.DBMigrate(db)

		container = &AppContainer{
			DB:     db,
			Router: router,
		}
	}
	return container
}

func GetDBConnection() *gorm.DB {
	if container != nil {
		return container.DB
	}
	return nil
}
