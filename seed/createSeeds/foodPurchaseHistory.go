package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
)

// CreateSeedFoodPurchaseHistories ...ここで記入したデータをDBにinsertする
func CreateSeedFoodPurchaseHistories() {

	foodPurchaseHistoriesInfos := []map[string]string{
		map[string]string{
			"UserID":    "goya",
			"RecipeID":  "4",
			"FoodID":    "1",
			"FoodCount": "1",
			"Quantity":  "500",
			"Unit":      "g",
		},
	}

	for _, info := range foodPurchaseHistoriesInfos {
		recipeID, _ := strconv.Atoi(info["RecipeID"])
		foodID, _ := strconv.Atoi(info["FoodID"])
		foodCount, _ := strconv.Atoi(info["FoodCount"])
		quantity, _ := strconv.Atoi(info["Quantity"])
		createFoodPurchaseHistory(model.FoodPurchaseHistory{
			UserID:    info["UserID"],
			RecipeID:  uint(recipeID),
			FoodID:    uint(foodID),
			FoodCount: uint(foodCount),
			Quantity:  uint(quantity),
			Unit:      info["Unit"],
		})
	}
}

// createIngredient ...渡されたデータをserviceの関数へ渡し、DBにinsertさせる
func createFoodPurchaseHistory(foodPurchaseHistory model.FoodPurchaseHistory) {
	foodPurchaseHistory, err := service.CreateFoodPurchaseHistory(foodPurchaseHistory)
	if err != nil {
		panic(err)
	}
	fmt.Printf("foodPurchaseHistory created\n")
}
