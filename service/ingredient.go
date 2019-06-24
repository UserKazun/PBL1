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

// GetIngredientsByRecipeIDAndFoodID ...
func GetIngredientsByRecipeIDAndFoodID(recipeID uint, foodID uint) (*model.Ingredient, error) {
	ingredient := model.Ingredient{}

	err := db.Where("recipe_id = ? and food_id = ?", recipeID, foodID).Find(&ingredient).Error
	if err != nil {
		return nil, err
	}

	return &ingredient, nil
}
