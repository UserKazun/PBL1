package service

import (
	"github.com/PBL1/model"
)

// CreateTradeItem ...DBに与えられたデータをinsertする
func CreateTradeItem(tradeItem model.TradeItem) (model.TradeItem, error) {
	err := db.Create(&tradeItem).Error
	if err != nil {
		return model.TradeItem{}, err
	}
	return tradeItem, nil
}

func GetTradeItems() ([]model.TradeItem, error) {
	tradeItems := []model.TradeItem{}

	err := db.Find(&tradeItems).Order("point", false).Error
	if err != nil {
		return nil, err
	}

	return tradeItems, nil
}

func GetTradeItemByID(id uint) (*model.TradeItem, error) {
	tradeItem := model.TradeItem{}

	err := db.Where("id = ?", id).First(&tradeItem).Error
	if err != nil {
		return nil, err
	}

	return &tradeItem, nil
}
