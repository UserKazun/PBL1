package service

import (
	"github.com/PBL1/model"
)

// CreateFoodPurchaseHistory ...DBに与えられたデータをinsertする
func CreateFoodPurchaseHistory(foodPurchaseHistory model.FoodPurchaseHistory) (model.FoodPurchaseHistory, error) {
	err := db.Create(&foodPurchaseHistory).Error
	if err != nil {
		return model.FoodPurchaseHistory{}, err
	}
	return foodPurchaseHistory, nil
}

func GetFoodPurchaseHistoriesByUserIDAndRecipeID(userID string, recipeID uint) ([]model.FoodPurchaseHistory, error) {
	foodPurchaseHistories := []model.FoodPurchaseHistory{}

	err := db.Where("user_id = ? and recipe_id = ?", userID, recipeID).Find(&foodPurchaseHistories).Error
	if err != nil {
		return nil, err
	}

	return foodPurchaseHistories, err

}
