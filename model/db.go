package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db = NewDBConn()

// NewDBConn ...DB接続
func NewDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	db.LogMode(true)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// GetDBConn ...
func GetDBConn() *gorm.DB {
	return db
}

// GetDBConfig ...DBに接続するための情報の取得
func GetDBConfig() (string, string) {
	return "mysql", "root:@tcp(127.0.0.1)/PBL1?charset=utf8&parseTime=True&loc=Local"
}
