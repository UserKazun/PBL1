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
	//PageURL     string
	Price uint
	Point uint
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

// RecipePurchaseHistory ...レシピの購入履歴
type RecipePurchaseHistory struct {
	ID          uint `gorm:"primary_key"` // 購入履歴のID
	UserID      string
	RecipeID    uint
	RecipeCount *uint
	Price       uint
	Point       uint
	CreatedAt   time.Time
}

// FoodPurchaseHistory ...食材の購入履歴
type FoodPurchaseHistory struct {
	RecipePurchaseHistoryID uint `sql:"type:int" gorm:"primary_key"`
	// UserID                  string
	// RecipeID                uint
	FoodID    uint `sql:"type:int" gorm:"primary_key"`
	FoodCount *uint
	Quantity  uint
	Unit      string
}

// DegreeOfContribution ...飢餓対策貢献ポイント
type DegreeOfContribution struct {
	UserID           string `gorm:"primary_key"`
	CumulativePoints uint
}

// Bookmark ...ブックマーク
type Bookmark struct {
	UserID   string `sql:"type:varchar(50)" gorm:"primary_key"`
	RecipeID uint   `sql:"type:int" gorm:"primary_key"`
}

// TradeItem ...交換アイテム
type TradeItem struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	ImageURL string
	Point    uint
}

// TradeItemHistory ...アイテム交換履歴
type TradeItemHistory struct {
	UserID      string `sql:"type:varchar(50)" gorm:"primary_key"`
	TradeItemID uint   `sql:"type:int" gorm:"primary_key"`
	CreatedAt   time.Time
}
