package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

// GetRecipeByMenuID ...受け取ったメニューIDを元にレシピのURLを返す
func GetRecipeByMenuID(c *gin.Context) {
	var recipeID uint
	var err error
	var recipePageURL string

	recipeID, err = GetUint(c, "recipe_id")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	recipePageURL, err = service.GetRecipeByMenuID(recipeID)
	if err != nil {
		log.Println("存在しないレシピIDです")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"recipe_PageURL": recipePageURL,
	})
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
		searchRecipe.PageURL = recipe.PageURL
		searchRecipe.Price = "￥" + strconv.FormatUint(uint64(recipe.Price), 10)

		searchRecipe.Point = recipe.Point

		searchRecipes = append(searchRecipes, searchRecipe)
	}

	c.JSON(http.StatusOK, searchRecipes)
}
