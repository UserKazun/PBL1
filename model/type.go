package model

import (
	"time"
)

// User ...ユーザー
type User struct {
	ID       int `sql:"type:int" gorm:"primary_key"`
	Password string
	Name     string
	IsAdmin  bool
}

// Menu ...メニュー
type Menu struct {
	ID   int `gorm:"primary_key"`
	Name string
}

// Food ...食材
type Food struct {
	ID             int `gorm:"primary_key"`
	Name           string
	Price          int
	ExpirationDate time.Time
	Point          int
}

// Material ...材料
type Material struct {
	MenuID int `sql:"type:int" gorm:"primary_key"`
	FoodID int `sql:"type:int" gorm:"primary_key"`
}

// Recipe ...レシピ
type Recipe struct {
	MenuID int `sql:"type:int" gorm:"primary_key"`
	URL    string
}
