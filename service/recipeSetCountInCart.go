package service

import (
	"github.com/PBL1/model"
)

// PutCartsRecipeCountByUserIDAndRecipeID ...
func PutCartsRecipeCountByUserIDAndRecipeID(userID string, recipeID uint, recipeCount uint) error {
	BeforeRecipeSetCountInCart := model.RecipeSetCountInCart{}
	var err error

	err = db.Where("user_id = ? and recipe_id = ?", userID, recipeID).First(&BeforeRecipeSetCountInCart).Error
	if err != nil {
		return err
	}

	afterRecipeSetCountInCart := BeforeRecipeSetCountInCart
	afterRecipeSetCountInCart.RecipeCount = recipeCount

	err = db.Model(&BeforeRecipeSetCountInCart).Update(afterRecipeSetCountInCart).Error
	if err != nil {
		return err
	}

	return nil
}