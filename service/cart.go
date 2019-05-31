package service

import (
	"github.com/PBL1/model"
)

// PostIngredientsToCart ...渡された材料データをカートテーブルに追加する
func PostIngredientsToCart(userID string, modelIngredients []model.Ingredient) error {
	cart := model.Cart{}

	for _, modelIngredient := range modelIngredients {
		cart.UserID = userID
		cart.RecipeID = modelIngredient.RecipeID
		cart.FoodID = modelIngredient.FoodID
		cart.Quantity = modelIngredient.Quantity
		cart.Unit = modelIngredient.Unit
		err := db.Create(&cart).Error
		if err != nil {
			return err
		}
	}

	return nil
}
