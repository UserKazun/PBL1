package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	db.LogMode(true)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDBConn() *gorm.DB {
	return db
}

func GetDBConfig() (string, string) {
	return "mysql", ""
}
