package model

// User ...ユーザー
type User struct {
	ID       string `gorm:"primary_key"`
	Password string
	Name     string
	IsAdmin  bool
}

// Recipe ...レシピ
type Recipe struct {
	ID         uint `gorm:"primary_key"`
	Name       string
	CategoryID uint
	URL        string
	Price      uint
	Point      uint
}

// Category ...料理カテゴリー
type Category struct {
	ID   uint   `gorm:"primary_key"`
	Name string //和食・洋食・イタリアン・中華・その他
}

// Food ...食材
type Food struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Point uint
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
	UserID   string `sql:"type:varchar(50)" gorm:"primary_key"`
	RecipeID uint   `sql:"type:int" gorm:"primary_key"`
	FoodID   uint   `sql:"type:int" gorm:"primary_key"`
	Quantity uint
	Unit     string
}
