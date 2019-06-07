package service

import (
	"time"

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

// GetRecipePurchaseHistoriesByUserID ...ユーザごとの購入履歴モデルデータを取得する
func GetRecipePurchaseHistoriesByUserID(userID string) ([]model.RecipePurchaseHistory, error) {
	recipePurchaseHistories := []model.RecipePurchaseHistory{}

	err := db.Where("user_id = ?", userID).Find(&recipePurchaseHistories).Error
	if err != nil {
		return nil, err
	}

	return recipePurchaseHistories, err
}

// GetPurchaseDatesByUserID ...ユーザごとの購入履歴データの購入日時群を取得
func GetPurchaseDatesByUserID(userID string) ([]time.Time, error) {
	recipePurchaseHistories := []model.RecipePurchaseHistory{}
	purchaseDates := []time.Time{}
	var err error

	err = db.Raw("select distinct created_at FROM recipe_purchase_histories where user_id = ? ORDER BY created_at desc LIMIT 5;", userID).Scan(&recipePurchaseHistories).Error
	if err != nil {
		return nil, err
	}

	for _, recipePurchaseHistory := range recipePurchaseHistories {
		purchaseDates = append(purchaseDates, recipePurchaseHistory.CreatedAt)
	}

	return purchaseDates, nil
}

func GetmodelRecipePurchaseHistoriesByUserIDAndPurchaseDate(userID string, purchaseDate time.Time) ([]model.RecipePurchaseHistory, error) {
	recipePurchaseHistories := []model.RecipePurchaseHistory{}

	err := db.Where("user_id = ? and created_at = ?", userID, purchaseDate).Find(&recipePurchaseHistories).Error
	if err != nil {
		return nil, err
	}

	return recipePurchaseHistories, err
}
