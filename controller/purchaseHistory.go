package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

// GetPurchaseHistoriesByUserID ...ユーザーIDを元に購入履歴を取得する
func GetPurchaseHistoriesByUserID(c *gin.Context) {
	purchaseHistory := PurchaseHistory{}
	purchaseHistories := []PurchaseHistory{}
	purchaseHistoryCard := PurchaseHistoryCard{}
	purchaseHistoryCards := []PurchaseHistoryCard{}
	foods := []Food{}
	food := Food{}
	modelRecipePurchaseHistories := []model.RecipePurchaseHistory{}
	modelFoodPurchaseHistories := []model.FoodPurchaseHistory{}
	purchaseDates := []time.Time{}
	recipe := model.Recipe{}
	var err error

	userID := c.Param("user_id")

	// errCode := AuthCheck(c, userID)
	// if errCode != nil {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	//特定ユーザーの購入日データ群を取得
	purchaseDates, err = service.GetPurchaseDatesByUserID(userID)
	if err != nil {
		log.Println("登録されていないユーザIDです")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for _, purchaseDate := range purchaseDates {
		purchaseHistory.Date = purchaseDate
		modelRecipePurchaseHistories, err = service.GetmodelRecipePurchaseHistoriesByUserIDAndPurchaseDate(userID, purchaseDate)

		for _, modelRecipePurchaseHistory := range modelRecipePurchaseHistories {
			// 材料の中身をリセット
			foods = nil

			// レシピの値段・ポイント・個数を挿入
			purchaseHistoryCard.Price = modelRecipePurchaseHistory.Price
			purchaseHistoryCard.Point = modelRecipePurchaseHistory.Point
			purchaseHistoryCard.RecipeCount = *modelRecipePurchaseHistory.RecipeCount

			// レシピデータを取得
			recipe = service.GetRecipeByRecipeID(modelRecipePurchaseHistory.RecipeID)

			// レシピの名前・イメージ画像URLを挿入
			purchaseHistoryCard.RecipeName = recipe.Name
			purchaseHistoryCard.RecipeImageURL = recipe.ImageURL

			// 値段とポイントの合計値を計算
			purchaseHistory.TotalPrice += purchaseHistoryCard.Price
			purchaseHistory.TotalPoint += purchaseHistoryCard.Point

			// レシピごとの食材購入履歴データを取得
			modelFoodPurchaseHistories, err = service.GetFoodPurchaseHistoriesByUserIDAndRecipeID(userID, recipe.ID)
			if err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}

			for _, modelFoodPurchaseHistory := range modelFoodPurchaseHistories {
				food.ID = modelFoodPurchaseHistory.FoodID
				food.FoodCount = *modelFoodPurchaseHistory.FoodCount
				food.Name, err = service.GetFoodNameByID(modelFoodPurchaseHistory.FoodID)

				if modelFoodPurchaseHistory.Quantity == 0 {
					food.Quantity = modelFoodPurchaseHistory.Unit
				} else {
					food.Quantity = strconv.FormatUint(uint64(modelFoodPurchaseHistory.Quantity), 10) + modelFoodPurchaseHistory.Unit
				}
				foods = append(foods, food)
			}

			purchaseHistoryCard.Foods = foods

			purchaseHistoryCards = append(purchaseHistoryCards, purchaseHistoryCard)
		}
		purchaseHistory.PurchaseHistoryCards = purchaseHistoryCards
	}
	purchaseHistories = append(purchaseHistories, purchaseHistory)

	c.JSON(http.StatusOK, purchaseHistories)

}
