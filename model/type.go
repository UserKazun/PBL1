package model

import (
	"time"
)

// User ...ユーザー
type User struct {
	ID       string `gorm:"primary_key"`
	Password string
	Name     string
	IsAdmin  bool
}

// Recipe ...レシピ
type Recipe struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Category uint
	URL      string
}

// Category ...料理カテゴリー
type Category struct {
	ID   uint   `gorm:"primary_key"`
	Name string //和食・洋食・イタリアン・中華・その為
}

// Food ...食材
type Food struct {
	ID             uint `gorm:"primary_key"`
	Name           string
	Price          uint
	ExpirationDate time.Time
	Point          uint
}

type Foods []Food

// ingredient ...材料
type Ingredient struct {
	MenuID uint `sql:"type:int" gorm:"primary_key"`
	FoodID uint `sql:"type:int" gorm:"primary_key"`
}

// Recipe ...レシピ
type Cart struct {
	UserID  string `gorm:"primary_key"`
	FoodSet FoodSet
}

type FoodSet struct {
	MenuName string
	Foods    Foods
}
