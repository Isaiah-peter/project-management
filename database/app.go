package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "isaiah:Etanuwoma18@/project-management?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
