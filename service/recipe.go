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

	return recipe.PageURL, err
}

// GetRecipeNameByRecipeID ...メニューIDを元に、レシピの名前を返す
func GetRecipeNameByRecipeID(recipeID uint) string {
	recipe := model.Recipe{}

	db.Where("id = ?", recipeID).First(&recipe)

	return recipe.Name
}

func GetRecipeByRecipeID(recipeID uint) model.Recipe {
	recipe := model.Recipe{}

	db.Where("id = ?", recipeID).First(&recipe)

	return recipe
}

//GetRecipesSearch ...
func GetRecipesSearch(categoryID uint, searchKey string) ([]model.Recipe, error) {
	recipes := []model.Recipe{}
	searchKey = "%" + searchKey + "%"

	if categoryID == 1 {
		err := db.Raw("select distinct recipes.id, recipes.name, recipes.description, recipes.image_url, recipes.page_url, recipes.price, recipes.point FROM recipes JOIN ingredients ON (recipes.ID=ingredients.recipe_id) LEFT JOIN foods ON (ingredients.food_id = foods.id) WHERE recipes.name LIKE ? OR foods.name LIKE ? ORDER BY RAND() LIMIT 5;", searchKey, searchKey).Scan(&recipes).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := db.Raw("select distinct recipes.id, recipes.name, recipes.description, recipes.image_url, recipes.page_url, recipes.price, recipes.point FROM recipes JOIN ingredients ON (recipes.ID=ingredients.recipe_id) LEFT JOIN foods ON (ingredients.food_id = foods.id) WHERE (recipes.name LIKE ? OR foods.name LIKE ?) AND recipes.category_id = ? ORDER BY RAND() LIMIT 5;", searchKey, searchKey, categoryID).Scan(&recipes).Error
		if err != nil {
			return nil, err
		}
	}

	return recipes, nil
}
