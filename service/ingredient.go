package service

import (
	"github.com/PBL1/model"
)

// CreateIngredient ...DBに与えられたデータをinsertする
func CreateIngredient(ingredient model.Ingredient) (model.Ingredient, error) {
	err := db.Create(&ingredient).Error
	if err != nil {
		return model.Ingredient{}, err
	}
	return ingredient, nil
}

// GetIngredientsByRecipeID ...レシピIDを元にそのレシピに必要な材料の情報を返す
func GetIngredientsByRecipeID(recipeID uint) []model.Ingredient {
	ingredients := []model.Ingredient{}

	db.Where("recipe_id = ?", recipeID).Find(&ingredients)

	return ingredients
}
