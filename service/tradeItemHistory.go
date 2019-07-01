package service

import (
	"log"
	"time"

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

func GetTradeDatesByUserID(userID string) ([]time.Time, error) {
	tradeItemHistories := []model.TradeItemHistory{}
	purchaseDates := []time.Time{}
	var err error

	err = db.Raw("select distinct created_at FROM trade_item_histories where user_id = ? ORDER BY created_at desc LIMIT 5;", userID).Scan(&tradeItemHistories).Error
	if err != nil {
		return nil, err
	}

	for _, tradeItemHistory := range tradeItemHistories {
		purchaseDates = append(purchaseDates, tradeItemHistory.CreatedAt)
	}

	return purchaseDates, nil
}

func GetTradeItemHistoriesByUserIDAndTradeDate(userID string, tradeDate string, tradeNextDate string) ([]model.TradeItemHistory, error) {
	tradeItemHistories := []model.TradeItemHistory{}

	log.Println("今日", tradeDate)
	log.Println("明日", tradeNextDate)

	err := db.Where("user_id = ? and created_at >= ? and created_at < ?", userID, tradeDate, tradeNextDate).Find(&tradeItemHistories).Error
	if err != nil {
		return nil, err
	}

	log.Println(tradeItemHistories)

	return tradeItemHistories, err
}
