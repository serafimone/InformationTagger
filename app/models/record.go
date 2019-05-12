package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Record type represents records
type Record struct {
	gorm.Model
	Content    string `json:"content"`
	DocumentID uint   `json:"document_id"`
}
