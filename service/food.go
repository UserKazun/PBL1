package service

import (
	"github.com/PBL1/model"
)

// CreateFood ...DBに与えられたデータをinsertする
func CreateFood(food model.Food) (model.Food, error) {
	err := db.Create(&food).Error
	if err != nil {
		return model.Food{}, err
	}
	return food, nil
}

// GetFoodNameByID ...食材IDを元に食材の名前を返す
func GetFoodNameByID(foodID uint) (string, error) {
	food := model.Food{}

	err := db.Where("id = ?", foodID).Find(&food).Error
	if err != nil {
		return "", err
	}
	return food.Name, nil
}
