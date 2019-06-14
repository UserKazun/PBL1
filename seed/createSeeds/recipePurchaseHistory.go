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
			"Price":       "50",
			"Point":       "5",
		},
		map[string]string{
			"UserID":      "goya",
			"RecipeID":    "2",
			"RecipeCount": "2",
			"Price":       "300",
			"Point":       "30",
		},
	}

	for _, info := range recipePurchaseHistoriesInfos {
		recipeID, _ := strconv.Atoi(info["RecipeID"])
		recipeCount, _ := strconv.Atoi(info["RecipeCount"])
		price, _ := strconv.Atoi(info["Price"])
		point, _ := strconv.Atoi(info["Point"])
		uintRecipeCount := uint(recipeCount)
		createRecipePurchaseHistory(model.RecipePurchaseHistory{
			UserID:      info["UserID"],
			RecipeID:    uint(recipeID),
			RecipeCount: &uintRecipeCount,
			Price:       uint(price),
			Point:       uint(point),
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
