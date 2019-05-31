package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
)

// CreateSeedIngredients ...ここで記入したデータをDBにinsertする
func CreateSeedIngredients() {

	ingredientsInfos := []map[string]string{
		map[string]string{
			"RecipeID": "1",
			"FoodID":   "1",
			"Quantity": "250",
			"Unit":     "g",
		},
		map[string]string{
			"RecipeID": "1",
			"FoodID":   "2",
			"Quantity": "85",
			"Unit":     "g",
		},
		map[string]string{
			"RecipeID": "1",
			"FoodID":   "3",
			"Quantity": "",
			"Unit":     "適量",
		},
		map[string]string{
			"RecipeID": "2",
			"FoodID":   "1",
			"Quantity": "250",
			"Unit":     "g",
		},
		map[string]string{
			"RecipeID": "2",
			"FoodID":   "2",
			"Quantity": "100",
			"Unit":     "g",
		},
		map[string]string{
			"RecipeID": "2",
			"FoodID":   "4",
			"Quantity": "",
			"Unit":     "欠片4分の1ほど",
		},
	}

	for _, info := range ingredientsInfos {
		recipeID, _ := strconv.Atoi(info["RecipeID"])
		foodID, _ := strconv.Atoi(info["FoodID"])
		quantity, _ := strconv.Atoi(info["Quantity"])
		createIngredient(model.Ingredient{
			RecipeID: uint(recipeID),
			FoodID:   uint(foodID),
			Quantity: uint(quantity),
			Unit:     info["Unit"],
		})
	}
}

// createIngredient ...渡されたデータをserviceの関数へ渡し、DBにinsertさせる
func createIngredient(ingredient model.Ingredient) {
	ingredient, err := service.CreateIngredient(ingredient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ingredient created\n")
}
