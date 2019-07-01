package controller

// LoginUser ...ログインユーザ情報
type LoginUser struct {
	ID      string `json:"user_id"`
	Name    string `json:"user_name"`
	IsAdmin bool   `json:"is_admin"`
}

// Recipe ...レシピ
type Recipe struct {
	RecipeName  string        `json:"recipe_name"`
	ImageURL    string        `json:"recipe_image_url"`
	Ingredients []Ingredient2 `json:"recipe_Ingredients"`
}

// Ingredient2 ...レシピとその材料群
type Ingredient2 struct {
	FoodName string `json:"food_name"`
	Quantity string `json:"quantity"`
}

// Ingredient ...レシピとその材料群
type Ingredient struct {
	RecipeName      string          `json:"recipe_name"`
	IngredientCards IngredientCards `json:"Ingredient"`
}

// IngredientCard ...材料
type IngredientCard struct {
	FoodName string `json:"food_name"`
	Quantity string `json:"quantity"`
}

// IngredientCards ...材料群
type IngredientCards []IngredientCard

// Category ...カテゴリ
type Category struct {
	Name string `json:"category_name"`
}

// SearchRecipe ...レシピの検索結果
type SearchRecipe struct {
	ID          uint   `json:"recipe_id"`
	Name        string `json:"recipe_name"`
	Description string `json:"recipe_Description"`
	ImageURL    string `json:"recipe_image_url"`
	//PageURL     string `json:"recipe_page_url"`
	Price       string `json:"price"`
	Point       uint   `json:"point"`
}

// Cart ...カート
type Cart struct {
	RecipeID       uint   `json:"recipe_id"`
	RecipeName     string `json:"recipe_name"`
	RecipeCount    uint   `json:"recipe_count"`
	RecipeImageURL string `json:"recipe_image_url"`
	Price          uint   `json:"price"`
	Point          uint   `json:"point"`
	Foods          []Food `json:"foods"`
}

// PurchaseHistory ...購入履歴
type PurchaseHistory struct {
	Date                 string `json:"date"`
	TotalPrice           uint   `json:"total_price"`
	TotalPoint           uint   `json:"total_point"`
	PurchaseHistoryCards []PurchaseHistoryCard
}

// PurchaseHistoryCard ...
type PurchaseHistoryCard struct {
	RecipeName     string `json:"recipe_name"`
	RecipeCount    uint   `json:"recipe_count"`
	RecipeImageURL string `json:"recipe_image_url"`
	Price          uint   `json:"price"`
	Point          uint   `json:"point"`
	Foods          []Food `json:"foods"`
}

//Food ...食料
type Food struct {
	ID        uint   `json:"food_id"`
	Name      string `json:"food_name"`
	Quantity  string `json:"food_quantity"`
	FoodCount uint   `json:"food_count"`
}

//Mypage ...マイページ
type Mypage struct {
	UserID            string `json:"user_id"`
	UserName          string `json:"user_name"`
	UserEmail         string `json:"user_email"`
	UserPassword      string `json:"user_password"`
	UserStreetAddress string `json:"user_street_address"`
	CumulativePoints  uint   `json:"cumulative_points"`
}

// Bookmark ...ブックマーク
type Bookmark struct {
	ID          uint   `json:"recipe_id"`
	Name        string `json:"recipe_name"`
	Description string `json:"recipe_Description"`
	ImageURL    string `json:"recipe_image_url"`
	Price       string `json:"price"`
	Point       uint   `json:"point"`
}

// TradeItem ...交換アイテム
type TradeItem struct {
	Name     string `json:"trade_name"`
	ImageURL string `json:"trade_image_url"`
	Point    uint   `json:"trade_point"`
}

// TradeItemHistory ...交換アイテム
type TradeItemHistory struct {
	Date                  string                 `json:"date"`
	TradeItemHistoryCards []TradeItemHistoryCard `json:"trade_items"`
}

type TradeItemHistoryCard struct {
	Name     string `json:"trade_name"`
	ImageURL string `json:"trade_image_url"`
	Point    uint   `json:"trade_point"`
}
