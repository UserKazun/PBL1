package service

import (
	"github.com/PBL1/model"
)

// CreateTradeItemsByUserID ...
func CreateTradeItemsByUserID(userID string, tradeItemID uint) error {

	tradeItemHistory := model.TradeItemHistory{}

	tradeItemHistory.UserID = userID
	tradeItemHistory.TradeItemID = tradeItemID

	err := db.Create(&tradeItemHistory).Error
	if err != nil {
		return err
	}
	return nil
}
