package controller

// LoginUser ...ログインユーザ情報
type LoginUser struct {
	ID      string `json:"user_id"`
	Name    string `json:"user_name"`
	IsAdmin bool   `json:"is_admin"`
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

type Category struct {
	Name string `json:"category_name"`
}

type Categories []Category
