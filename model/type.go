package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User ...ユーザー
type User struct {
	ID       string `gorm:"primary_key"`
	Password string
	Name     string
	IsAdmin  bool
}

// Menu ...メニュー
type Menu struct {
	gorm.Model
	Name string
}

// Food ...食材
type Food struct {
	gorm.Model
	Name           string
	Price          uint
	ExpirationDate time.Time
	Point          uint
}

// Material ...材料
type Material struct {
	MenuID uint `sql:"type:int" gorm:"primary_key"`
	FoodID uint `sql:"type:int" gorm:"primary_key"`
}

// Recipe ...レシピ
type Recipe struct {
	MenuID uint `sql:"type:int" gorm:"primary_key"`
	URL    string
}
