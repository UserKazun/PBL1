package service

import (
	"github.com/PBL1/model"
)

// PutCartsFoodCount ...カート内の食料の数に関するテーブルを更新
func PutCartsFoodCount(userID string, recipeID uint, foodID uint, foodCount uint) error {
	BeforeCart := model.Cart{}
	var err error

	err = db.Where("user_id = ? and recipe_id = ? and food_id = ?", userID, recipeID, foodID).First(&BeforeCart).Error
	if err != nil {
		return err
	}

	afterCart := BeforeCart
	afterCart.FoodCount = &foodCount

	err = db.Model(&BeforeCart).Update(afterCart).Error
	if err != nil {
		return err
	}

	return nil
}

func InsertFoodCartContentsToPuchaseHistory(userID string, foodIDsInCarts []model.Cart, ingredientInUserCarts *model.Ingredient) error {
	foodPurchaseHistory := model.FoodPurchaseHistory{}
	ingredientInUserCart := ingredientInUserCarts

	// 引数で受け取ったfoodID分だけloop
	for _, foodIDsInCart := range foodIDsInCarts {
		// 該当する部分にデータを代入
		foodPurchaseHistory.UserID = userID
		foodPurchaseHistory.RecipeID = foodIDsInCart.RecipeID
		foodPurchaseHistory.FoodID = foodIDsInCart.FoodID
		foodPurchaseHistory.FoodCount = foodIDsInCart.FoodCount
		foodPurchaseHistory.Quantity = foodIDsInCart.FoodID * ingredientInUserCart.Quantity
		foodPurchaseHistory.Unit = ingredientInUserCart.Unit

		_, err := CreateFoodPurchaseHistory(foodPurchaseHistory)

		return err
	}

	return nil
}
