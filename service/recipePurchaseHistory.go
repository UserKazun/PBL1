package service

import (
	"github.com/PBL1/model"
)

// CreateRecipePurchaseHistory ...DBに与えられたデータをinsertする
func CreateRecipePurchaseHistory(recipePurchaseHistory model.RecipePurchaseHistory) (model.RecipePurchaseHistory, error) {
	err := db.Create(&recipePurchaseHistory).Error
	if err != nil {
		return model.RecipePurchaseHistory{}, err
	}
	return recipePurchaseHistory, nil
}
