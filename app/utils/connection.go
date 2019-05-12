package utils

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/serafimone/InformationTagger/config"
)

func getConnectionString(config *config.DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset)
}

//GetConnection function getting db connection instance based on config data
func GetConnection(config *config.DBConfig) *gorm.DB {
	db, err := gorm.Open(config.Dialect, getConnectionString(config))
	if err != nil {
		log.Fatalf("[Error] %s", err)
	}
	return db
}
