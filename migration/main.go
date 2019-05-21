package main

import (
	"github.com/PBL1/model"
	_ "github.com/jinzhu/gorm"
)

func main() {

	db := model.GetDBConn()

	// migrate
	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.Recipe{})
	db.DropTableIfExists(&model.Food{})
	db.DropTableIfExists(&model.Ingredient{})
	db.DropTableIfExists(&model.Category{})

	db.CreateTable(&model.User{})
	db.CreateTable(&model.Recipe{})
	db.CreateTable(&model.Food{})
	db.CreateTable(&model.Category{})
	db.CreateTable(&model.Ingredient{})
}
