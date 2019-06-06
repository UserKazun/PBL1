package controller

// LoginUser ...ログインユーザ情報
type LoginUser struct {
	ID      string `json:"user_id"`
	Name    string `json:"user_name"`
	IsAdmin bool   `json:"is_admin"`
}

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
	PageURL     string `json:"recipe_page_url"`
	Price       string `json:"price"`
	Point       uint   `json:"point"`
}

// Cart ...カート
type Cart struct {
	RecipeName     string       `json:"recipe_name"`
	RecipeCount    uint         `json:"recipe_count"`
	RecipeImageURL string       `json:"recipe_image_url"`
	Price          uint         `json:"price"`
	Point          uint         `json:"point"`
	FoodsInCart    []FoodInCart `json:"food_names"`
}

//FoodInCart ...カート内の食料
type FoodInCart struct {
	Name      string `json:"food_name"`
	FoodCount uint   `json:"food_count"`
	Quantity  string `json:"food_quantity"`
}
