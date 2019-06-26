package controller

import (
	"log"
	"net/http"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

func GetTradeItemsByUserID(c *gin.Context) {
	tradeItems := []model.TradeItem{}

	tradeItems, err := service.GetTradeItemsByUserID()
	if err != nil {
		log.Println("交換商品データが取得できませんでした")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tradeItems)
}
