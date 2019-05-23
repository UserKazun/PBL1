package service

import (
	"github.com/PBL1/model"
)

// CreateFood ...DBに与えられたデータをinsertする
// func CreateFood(food model.Food) (model.Food, error) {
// 	err := db.Create(&food).Error
// 	if err != nil {
// 		return model.Food{}, err
// 	}
// 	return food, nil
// }

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
