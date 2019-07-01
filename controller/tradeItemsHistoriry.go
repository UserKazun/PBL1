package controller

import (
	"net/http"
	"time"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

func GetTradeItemsHistoriesByUserID(c *gin.Context) {
	tradeDates := []time.Time{}
	timeNext := time.Time{}
	tradeItemHistory := TradeItemHistory{}
	tradeItemHistories := []TradeItemHistory{}
	modelTradeItemHistories := []model.TradeItemHistory{}
	tradeItemHistoryCard := TradeItemHistoryCard{}
	tradeItemHistoryCards := []TradeItemHistoryCard{}
	tradeDateTemp := "0"
	var tradeNextDate string
	var err error

	userID := c.Param("user_id")
	// model.TradeItemHistory

	//特定ユーザーの交換日データ群を取得
	tradeDates, err = service.GetTradeDatesByUserID(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, tradeDate := range tradeDates {
		// すでに挿入された日付か確認
		if tradeDate.Format("2006-01-02") == tradeDateTemp {
			continue
		}

		// 挿入された日付を保持
		tradeDateTemp = tradeDate.Format("2006-01-02")

		tradeItemHistory.Date = tradeDate.Format("2006-01-02")

		timeNext = tradeDate.Add(24 * time.Hour)
		tradeNextDate = timeNext.Format("2006-01-02")

		// 指定された日に交換された商品を取得
		modelTradeItemHistories, err = service.GetTradeItemHistoriesByUserIDAndTradeDate(userID, tradeItemHistory.Date, tradeNextDate)

		tradeItemHistoryCards = []TradeItemHistoryCard{}
		// 取得した商品をその日の購入履歴データとして挿入
		for _, modelTradeItemHistory := range modelTradeItemHistories {
			tradeItem, err := service.GetTradeItemByID(modelTradeItemHistory.TradeItemID)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			tradeItemHistoryCard.Name = tradeItem.Name
			tradeItemHistoryCard.ImageURL = tradeItem.ImageURL
			tradeItemHistoryCard.Point = tradeItem.Point

			tradeItemHistoryCards = append(tradeItemHistoryCards, tradeItemHistoryCard)
		}
		tradeItemHistory.TradeItemHistoryCards = tradeItemHistoryCards

		tradeItemHistories = append(tradeItemHistories, tradeItemHistory)
	}

	c.JSON(http.StatusOK, tradeItemHistories)
}
