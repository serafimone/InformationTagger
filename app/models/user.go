package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UUID     string `json:"uuid" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
