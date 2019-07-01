package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

func GetTradeItems(c *gin.Context) {
	tradeItems := []model.TradeItem{}

	tradeItems, err := service.GetTradeItems()
	if err != nil {
		log.Println("交換商品データが取得できませんでした")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tradeItems)
}

func PostTradeItemsByUserID(c *gin.Context) {
	// tradeItems := []model.TradeItem{}
	var err error
	moneyFlag := false

	userID := c.PostForm("user_id")
	tradeItemID, _ := strconv.Atoi(c.PostForm("trade_item_id"))

	err, moneyFlag = service.ReduceCumulativePointsByUserID(userID, uint(tradeItemID))
	if moneyFlag {
		log.Println("手持ちの飢餓ポイントが足りませんでした")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Println("合計飢餓対策貢献ポイントを減少できませんでした")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = service.CreateTradeItemsByUserID(userID, uint(tradeItemID))
	if err != nil {
		log.Println("商品を交換できませんでした")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
