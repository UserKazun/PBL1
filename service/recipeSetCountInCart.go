package service

import (
	"log"

	"github.com/PBL1/model"
)

// PutCartsRecipeCount ...カート内のレシピセット数に関するテーブルを更新
func PutCartsRecipeCount(userID string, recipeID uint, recipeCount uint) error {
	beforeRecipeSetCountInCart := model.RecipeSetCountInCart{}
	var err error

	err = db.Where("user_id = ? and recipe_id = ?", userID, recipeID).First(&beforeRecipeSetCountInCart).Error
	if err != nil {
		return err
	}

	afterRecipeSetCountInCart := beforeRecipeSetCountInCart
	afterRecipeSetCountInCart.RecipeCount = &recipeCount

	err = db.Model(&beforeRecipeSetCountInCart).Update(afterRecipeSetCountInCart).Error
	if err != nil {
		return err
	}

	log.Print(afterRecipeSetCountInCart.RecipeCount)

	return nil
}

func GetRecipeSetCountInCartsByUserID(userID string) ([]model.RecipeSetCountInCart, error){
	recipeSetCountInCarts := []model.RecipeSetCountInCart{}

	err := db.Where("user_id = ?", userID).Find(&recipeSetCountInCarts).Error
	if err != nil {
		return nil, err
	}
	return recipeSetCountInCarts, nil
}

func GetRecipePriceAndPointByID(recipeID uint) (*uint, *uint, error){
	recipe := model.Recipe{}

	err := db.Where("id = ?", recipeID).First(&recipe).Error
	if err != nil {
		return nil, nil, err
	}
	return &recipe.Price, &recipe.Point, nil
}
