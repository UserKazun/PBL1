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
