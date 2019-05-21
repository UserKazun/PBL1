package service

import (
	"github.com/PBL1/model"
)

// GetRecipeByMenuID ...メニューIDを受け取り、レシピのURLを返す
func GetRecipeByMenuID(recipeID uint) (string, error) {
	recipe := model.Recipe{}

	err := db.Where("id = ?", recipeID).First(&recipe).Error

	return recipe.URL, err
}

// CreateRecipe ...DBに与えられたデータをinsertする
func CreateRecipe(recipe model.Recipe) (model.Recipe, error) {
	err := db.Create(&recipe).Error
	if err != nil {
		return model.Recipe{}, err
	}
	return recipe, nil
}
