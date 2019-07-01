package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

func GetRecipeByRecipeID(c *gin.Context) {
	recipe := Recipe{}

	ingredient := Ingredient2{}
	ingredients := []Ingredient2{}

	var err error

	recipeID, err := GetUint(c, "recipe_id")
	if err != nil {
		log.Println("存在しないレシピIDです")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	modelRecipe := service.GetRecipeByRecipeID(recipeID)

	modelIngredients := service.GetIngredientsByRecipeID(recipeID)

	for _, modelIngredient := range modelIngredients {
		ingredient.FoodName, err = service.GetFoodNameByID(modelIngredient.FoodID)
		if err != nil {
			log.Println("存在しない食材IDです")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if modelIngredient.Quantity == 0 {
			ingredient.Quantity = modelIngredient.Unit
		} else {
			ingredient.Quantity = strconv.FormatUint(uint64(modelIngredient.Quantity), 10) + modelIngredient.Unit
		}
		ingredients = append(ingredients, ingredient)
	}

	recipe.RecipeName = modelRecipe.Name
	recipe.ImageURL = modelRecipe.ImageURL
	recipe.Ingredients = ingredients

	c.JSON(http.StatusOK, recipe)
}

//GetRecipesSearch ...与えられたキーを元に検索した結果のレシピデータを取得する
func GetRecipesSearch(c *gin.Context) {
	searchRecipe := SearchRecipe{}
	searchRecipes := []SearchRecipe{}
	var categoryID uint
	var searchKey string
	var err error

	categoryID, err = GetUint(c, "category_id")
	if err != nil {
		log.Println("category_id：数値が入力されていません")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	searchKey = c.Param("search_key")

	recipes, err := service.GetRecipesSearch(categoryID, searchKey)
	if err != nil {
		log.Println("与えられた検索キーが不適切です")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for _, recipe := range recipes {
		searchRecipe.ID = recipe.ID
		searchRecipe.Name = recipe.Name
		searchRecipe.Description = recipe.Description
		searchRecipe.ImageURL = recipe.ImageURL
		searchRecipe.Price = strconv.FormatUint(uint64(recipe.Price), 10) + "円"

		searchRecipe.Point = recipe.Point

		searchRecipes = append(searchRecipes, searchRecipe)
	}

	c.JSON(http.StatusOK, searchRecipes)
}

func GetRecipesCategoryOnlySearch(c *gin.Context) {
	searchRecipe := SearchRecipe{}
	searchRecipes := []SearchRecipe{}
	var categoryID uint
	var err error

	categoryID, err = GetUint(c, "category_id")
	if err != nil {
		log.Println("category_id：数値が入力されていません")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	recipes, err := service.GetRecipesCategoryOnlySearch(categoryID)
	if err != nil {
		log.Println("与えられた検索キーが不適切です")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for _, recipe := range recipes {
		searchRecipe.ID = recipe.ID
		searchRecipe.Name = recipe.Name
		searchRecipe.Description = recipe.Description
		searchRecipe.ImageURL = recipe.ImageURL
		searchRecipe.Price = "¥" + strconv.FormatUint(uint64(recipe.Price), 10)

		searchRecipe.Point = recipe.Point

		searchRecipes = append(searchRecipes, searchRecipe)
	}

	c.JSON(http.StatusOK, searchRecipes)
}
