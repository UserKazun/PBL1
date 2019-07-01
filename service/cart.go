package service

import (
	"github.com/PBL1/model"
)

// PostIngredientsToCart ...渡された材料データをカートテーブルに追加する
func PostIngredientsToCart(userID string, modelIngredients []model.Ingredient) error {
	cart := model.Cart{}
	numOne := uint(1)

	for _, modelIngredient := range modelIngredients {
		cart.UserID = userID
		cart.RecipeID = modelIngredient.RecipeID
		cart.FoodID = modelIngredient.FoodID
		cart.FoodCount = &numOne
		err := db.Create(&cart).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// GetRecipeIDsInCartByUserID ...ユーザーごとのカート内のレシピIDを返す
func GetRecipeIDsInCartByUserID(userID string) ([]uint, error) {
	carts := []model.Cart{}
	recipeIDs := []uint{}

	err := db.Raw("select distinct carts.recipe_id FROM carts WHERE user_id = ?;", userID).Scan(&carts).Error

	if err != nil {
		return nil, err
	}

	for _, cart := range carts {
		recipeIDs = append(recipeIDs, cart.RecipeID)
	}

	return recipeIDs, nil
}

// InsertRecipeCount ...ユーザーごとのカートにあるレシピのセット数（○人前）をTableにInsertする
func InsertRecipeCount(userID string, recipeID uint) error {
	recipeSetCountInCart := model.RecipeSetCountInCart{}
	numOne := uint(1)

	recipeSetCountInCart.UserID = userID
	recipeSetCountInCart.RecipeID = recipeID
	recipeSetCountInCart.RecipeCount = &numOne
	err := db.Create(&recipeSetCountInCart).Error
	if err != nil {
		return err
	}

	return nil
}

// GetFoodIDsInCartByUserID ...
func GetFoodIDsInCartByUserID(userID string, recipeID uint) ([]model.Cart, error) {
	carts := []model.Cart{}

	err := db.Where("user_id = ? and recipe_id = ?", userID, recipeID).Find(&carts).Error
	if err != nil {
		return nil, err
	}

	return carts, nil
}

// DeleteCartContent ...カートの中身を削除する
func DeleteCartContent(userID string) error {
	carts := []model.Cart{}

	err := db.Delete(&carts).Error
	if err != nil {
		return err
	}

	return nil
}
