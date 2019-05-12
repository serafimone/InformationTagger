package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Document{}, &Record{})
	db.Model(&Record{}).AddForeignKey("document_id", "documents(id)", "CASCADE", "CASCADE")
	return db
}
