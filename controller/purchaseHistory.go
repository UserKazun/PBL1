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

func GetPurchaseHistoriesByUserID(c *gin.Context) {
	purchaseHistory := PurchaseHistory{}
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

	errCode := AuthCheck(c, userID)
	if errCode != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	modelRecipePurchaseHistories, err = service.GetRecipePurchaseHistoriesByUserID(userID)
	if err != nil {
		log.Println("登録されていないユーザIDです")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for _, modelRecipePurchaseHistory := range modelRecipePurchaseHistories {
		purchaseHistory.Date = modelRecipePurchaseHistory.CreatedAt
		purchaseDates, err = service.GetPurchaseDatesByUserID(userID)

		for _, purchaseDate := range purchaseDates {
			modelRecipePurchaseHistories, err = service.GetmodelRecipePurchaseHistoriesByUserIDAndPurchaseDate(userID, purchaseDate)

			for _, modelRecipePurchaseHistory := range modelRecipePurchaseHistories {
				purchaseHistoryCard.Point = modelRecipePurchaseHistory.Point
				purchaseHistoryCard.Price = modelRecipePurchaseHistory.Price
				purchaseHistoryCard.RecipeCount = modelRecipePurchaseHistory.RecipeCount

				recipe = service.GetRecipeByRecipeID(modelRecipePurchaseHistory.RecipeID)

				purchaseHistoryCard.RecipeName = recipe.Name
				purchaseHistoryCard.RecipeImageURL = recipe.ImageURL

				purchaseHistory.TotalPrice += purchaseHistoryCard.Price
				purchaseHistory.TotalPoint += purchaseHistoryCard.Point

				modelFoodPurchaseHistories, err = service.GetFoodPurchaseHistoriesByUserIDAndRecipeID(userID, recipe.ID)
				if err != nil {
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}

				for _, modelFoodPurchaseHistory := range modelFoodPurchaseHistories {
					food.FoodCount = modelFoodPurchaseHistory.FoodCount
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

	}

	c.JSON(http.StatusOK, purchaseHistory)

}
