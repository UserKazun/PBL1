package model

import (
	"time"
)

// User ...ユーザー
type User struct {
	ID            string `gorm:"primary_key"`
	Password      string
	Name          string
	Email         string
	StreetAddress string
	IsAdmin       bool
}

// Recipe ...レシピ
type Recipe struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Description string
	CategoryID  uint
	ImageURL    string
	PageURL     string
	Price       uint
	Point       uint
}

// Category ...料理カテゴリー
type Category struct {
	ID   uint   `gorm:"primary_key"`
	Name string //和食・洋食・イタリアン・中華・その他
}

// Food ...食材
type Food struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

// Ingredient ...材料
type Ingredient struct {
	RecipeID uint `sql:"type:int" gorm:"primary_key"`
	FoodID   uint `sql:"type:int" gorm:"primary_key"`
	Quantity uint
	Unit     string
}

// Cart ...カート
type Cart struct {
	UserID    string `sql:"type:varchar(50)" gorm:"primary_key"`
	RecipeID  uint   `sql:"type:int" gorm:"primary_key"`
	FoodID    uint   `sql:"type:int" gorm:"primary_key"`
	FoodCount *uint
}

// RecipeSetCountInCart ...カート内のレシピセット数
type RecipeSetCountInCart struct {
	UserID      string `sql:"type:varchar(50)" gorm:"primary_key"`
	RecipeID    uint   `sql:"type:int" gorm:"primary_key"`
	RecipeCount *uint
}

// FoodPurchaseHistory ...食材の購入履歴
type FoodPurchaseHistory struct {
	UserID    string `sql:"type:varchar(50)" gorm:"primary_key"`
	RecipeID  uint   `sql:"type:int" gorm:"primary_key"`
	FoodID    uint   `sql:"type:int" gorm:"primary_key"`
	FoodCount *uint
	Quantity  uint
	Unit      string
}

// RecipePurchaseHistory ...レシピの購入履歴
type RecipePurchaseHistory struct {
	UserID      string `sql:"type:varchar(50)" gorm:"primary_key"`
	RecipeID    uint   `sql:"type:int" gorm:"primary_key"`
	RecipeCount *uint
	Price       uint
	Point       uint
	CreatedAt   time.Time
}

// DegreeOfContribution ...飢餓対策貢献ポイント
type DegreeOfContribution struct {
	UserID           string `gorm:"primary_key"`
	CumulativePoints uint
}

//Bookmark ...ブックマーク
type Bookmark struct {
	UserID   string `sql:"type:varchar(50)" gorm:"primary_key"`
	RecipeID uint   `sql:"type:int" gorm:"primary_key"`
}
