package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
)

// CreateSeedRecipePurchaseHistories ...ここで記入したデータをDBにinsertする
func CreateSeedRecipePurchaseHistories() {

	recipePurchaseHistoriesInfos := []map[string]string{
		map[string]string{
			"UserID":      "goya",
			"RecipeID":    "4",
			"RecipeCount": "1",
		},
	}

	for _, info := range recipePurchaseHistoriesInfos {
		recipeID, _ := strconv.Atoi(info["RecipeID"])
		recipeCount, _ := strconv.Atoi(info["RecipeCount"])
		createRecipePurchaseHistory(model.RecipePurchaseHistory{
			UserID:      info["UserID"],
			RecipeID:    uint(recipeID),
			RecipeCount: uint(recipeCount),
		})
	}
}

// createIngredient ...渡されたデータをserviceの関数へ渡し、DBにinsertさせる
func createRecipePurchaseHistory(recipePurchaseHistory model.RecipePurchaseHistory) {
	recipePurchaseHistory, err := service.CreateRecipePurchaseHistory(recipePurchaseHistory)
	if err != nil {
		panic(err)
	}
	fmt.Printf("recipePurchaseHistory created\n")
}
