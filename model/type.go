package model

import "github.com/jinzhu/gorm"

// ユーザー
type User struct {
	gorm.Model
	ID       int
	Password string
	Admin    bool
}

// メニュー
type Menu struct {
	gorm.Model
	ID       int
	MenuName string
}

// 食材
type Food struct {
	gorm.Model
	ID             int
	FoodName       string
	Price          int
	ExpirationDate string
}

// 食材
type Material struct {
	gorm.Model
	FoodId int
	MenuId int
}

// レシピURL
type RecipeUrl struct {
	gorm.Model
	MenuId int
	Url    string
}
