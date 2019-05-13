package main

import (
	"github.com/PBL1/model"
	_ "github.com/jinzhu/gorm"
)

func main() {

	db := model.GetDBConn()

	// migrate
	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.Menu{})
	db.DropTableIfExists(&model.Food{})
	db.DropTableIfExists(&model.Material{})
	db.DropTableIfExists(&model.Recipe{})

	db.CreateTable(&model.User{})
	db.CreateTable(&model.Menu{})
	db.CreateTable(&model.Food{})
	db.CreateTable(&model.Material{})
	db.CreateTable(&model.Recipe{})

}
