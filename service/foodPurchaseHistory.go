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

func GetFoodPurchaseHistoriesByUserIDAndRecipeID(recipePurchaseHistoryID uint) ([]model.FoodPurchaseHistory, error) {
	foodPurchaseHistories := []model.FoodPurchaseHistory{}

	err := db.Where("recipe_purchase_history_id = ?", recipePurchaseHistoryID).Find(&foodPurchaseHistories).Error
	if err != nil {
		return nil, err
	}

	return foodPurchaseHistories, err

}
