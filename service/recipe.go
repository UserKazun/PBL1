package service

import (
	"github.com/PBL1/model"
)

// CreateRecipe ...DBに与えられたデータをinsertする
func CreateRecipe(recipe model.Recipe) (model.Recipe, error) {
	err := db.Create(&recipe).Error
	if err != nil {
		return model.Recipe{}, err
	}
	return recipe, nil
}

// GetRecipeByMenuID ...メニューIDを元に、レシピのURLを返す
func GetRecipeByMenuID(recipeID uint) (string, error) {
	recipe := model.Recipe{}

	err := db.Where("id = ?", recipeID).First(&recipe).Error

	return recipe.URL, err
}

// GetRecipeNameByRecipeID ...メニューIDを元に、レシピの名前を返す
func GetRecipeNameByRecipeID(recipeID uint) (string, error) {
	recipe := model.Recipe{}

	err := db.Where("id = ?", recipeID).First(&recipe).Error

	return recipe.Name, err
}
