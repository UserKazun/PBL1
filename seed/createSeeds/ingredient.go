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
		map[string]string{
			"RecipeID": "4",
			"FoodID":   "1",
			"Quantity": "500",
			"Unit":     "g",
		},
		map[string]string{
			"RecipeID": "5",
			"FoodID":   "8",
			"Quantity": "1",
			"Unit":     "個",
		},

		map[string]string{
			"RecipeID": "5",
			"FoodID":   "9",
			"Quantity": "90",
			"Unit":     "g",
		},
		map[string]string{
			"RecipeID": "5",
			"FoodID":   "10",
			"Quantity": "",
			"Unit":     "お好み",
		},
		map[string]string{
			"RecipeID": "5",
			"FoodID":   "11",
			"Quantity": "",
			"Unit":     "1/2片",
		},
		map[string]string{
			"RecipeID": "5",
			"FoodID":   "12",
			"Quantity": "15",
			"Unit":     "g",
		},
		map[string]string{
			"RecipeID": "5",
			"FoodID":   "13",
			"Quantity": "",
			"Unit":     "たっぷり",
		},
		map[string]string{
			"RecipeID": "5",
			"FoodID":   "14",
			"Quantity": "",
			"Unit":     "適量",
		},
		map[string]string{
			"RecipeID": "6",
			"FoodID":   "15",
			"Quantity": "1",
			"Unit":     "丁",
		},
		map[string]string{
			"RecipeID": "6",
			"FoodID":   "16",
			"Quantity": "100",
			"Unit":     "g",
		},
		map[string]string{
			"RecipeID": "6",
			"FoodID":   "6",
			"Quantity": "",
			"Unit":     "1/4玉",
		},
		map[string]string{
			"RecipeID": "6",
			"FoodID":   "17",
			"Quantity": "",
			"Unit":     "大さじ1",
		},
		map[string]string{
			"RecipeID": "6",
			"FoodID":   "18",
			"Quantity": "",
			"Unit":     "大さじ2",
		},
		map[string]string{
			"RecipeID": "6",
			"FoodID":   "19",
			"Quantity": "",
			"Unit":     "少々",
		},
		map[string]string{
			"RecipeID": "7",
			"FoodID":   "21",
			"Quantity": "5",
			"Unit":     "匹",
		},
		map[string]string{
			"RecipeID": "7",
			"FoodID":   "22",
			"Quantity": "",
		},
		map[string]string{
			"RecipeID": "7",
			"FoodID":   "23",
			"Quantity": "",
			"Unit":     "大さじ1",
		}, map[string]string{
			"RecipeID": "7",
			"FoodID":   "24",
			"Quantity": "",
			"Unit":     "適量",
		},
		map[string]string{
			"RecipeID": "7",
			"FoodID":   "25",
			"Quantity": "",
			"Unit":     "適量",
		}, map[string]string{
			"RecipeID": "7",
			"FoodID":   "26",
			"Quantity": "",
			"Unit":     "お好み",
		},
		map[string]string{
			"RecipeID": "7",
			"FoodID":   "27",
			"Quantity": "",
			"Unit":     "お好み",
		},
		map[string]string{
			"RecipeID": "7",
			"FoodID":   "28",
			"Quantity": "",
			"Unit":     "少々",
		},
		map[string]string{
			"RecipeID": "7",
			"FoodID":   "13",
			"Quantity": "",
			"Unit":     "少々",
		},
		map[string]string{
			"RecipeID": "8",
			"FoodID":   "29",
			"Quantity": "2",
			"Unit":     "切",
		},
		map[string]string{
			"RecipeID": "8",
			"FoodID":   "6",
			"Quantity": "",
			"Unit":     "大きめの1/2個",
		},
		map[string]string{
			"RecipeID": "8",
			"FoodID":   "30",
			"Quantity": "",
			"Unit":     "1/2パック",
		},
		map[string]string{
			"RecipeID": "8",
			"FoodID":   "22",
			"Quantity": "",
			"Unit":     "適量",
		},
		map[string]string{
			"RecipeID": "8",
			"FoodID":   "19",
			"Quantity": "",
			"Unit":     "適量",
		},
		map[string]string{
			"RecipeID": "8",
			"FoodID":   "31",
			"Quantity": "",
			"Unit":     "大さじ1",
		},
		map[string]string{
			"RecipeID": "8",
			"FoodID":   "33",
			"Quantity": "",
			"Unit":     "適量",
		},
		map[string]string{
			"RecipeID": "9",
			"FoodID":   "34",
			"Quantity": "",
			"Unit":     "1/2本",
		},
		map[string]string{
			"RecipeID": "9",
			"FoodID":   "35",
			"Quantity": "5",
			"Unit":     "枚",
		},
		map[string]string{
			"RecipeID": "9",
			"FoodID":   "36",
			"Quantity": "",
			"Unit":     "適量",
		},
		map[string]string{
			"RecipeID": "9",
			"FoodID":   "37",
			"Quantity": "",
			"Unit":     "1/2個",
		},
		map[string]string{
			"RecipeID": "9",
			"FoodID":   "15",
			"Quantity": "",
			"Unit":     "1/2丁",
		},
		map[string]string{
			"RecipeID": "9",
			"FoodID":   "17",
			"Quantity": "",
			"Unit":     "大さじ1",
		},
		map[string]string{
			"RecipeID": "9",
			"FoodID":   "28",
			"Quantity": "",
			"Unit":     "小さじ1",
		},
		map[string]string{
			"RecipeID": "9",
			"FoodID":   "39",
			"Quantity": "",
			"Unit":     "大さじ1",
		},
		map[string]string{
			"RecipeID": "9",
			"FoodID":   "41",
			"Quantity": "",
			"Unit":     "大さじ1",
		},
		map[string]string{
			"RecipeID": "9",
			"FoodID":   "42",
			"Quantity": "2",
			"Unit":     "㎝",
		},
		map[string]string{
			"RecipeID": "9",
			"FoodID":   "43",
			"Quantity": "",
			"Unit":     "大さじ1",
		},
		map[string]string{
			"RecipeID": "10",
			"FoodID":   "44",
			"Quantity": "1",
			"Unit":     "枚",
		},
		map[string]string{
			"RecipeID": "10",
			"FoodID":   "8",
			"Quantity": "4",
			"Unit":     "個",
		},
		map[string]string{
			"RecipeID": "10",
			"FoodID":   "45",
			"Quantity": "50",
			"Unit":     "cc",
		},
		map[string]string{
			"RecipeID": "10",
			"FoodID":   "17",
			"Quantity": "30",
			"Unit":     "cc",
		},
		map[string]string{
			"RecipeID": "10",
			"FoodID":   "46",
			"Quantity": "",
			"Unit":     "中さじ1",
		},
		map[string]string{
			"RecipeID": "10",
			"FoodID":   "47",
			"Quantity": "",
			"Unit":     "適量",
		},
		map[string]string{
			"RecipeID": "11",
			"FoodID":   "1",
			"Quantity": "",
			"Unit":     "3/4合",
		},
		map[string]string{
			"RecipeID": "11",
			"FoodID":   "8",
			"Quantity": "1",
			"Unit":     "個",
		},
		map[string]string{
			"RecipeID": "11",
			"FoodID":   "48",
			"Quantity": "",
			"Unit":     "お好み",
		},
		map[string]string{
			"RecipeID": "11",
			"FoodID":   "49",
			"Quantity": "",
			"Unit":     "お好み",
		},
		map[string]string{
			"RecipeID": "11",
			"FoodID":   "40",
			"Quantity": "",
			"Unit":     "小さじ1",
		},
		map[string]string{
			"RecipeID": "11",
			"FoodID":   "50",
			"Quantity": "",
			"Unit":     "お好み",
		},
		map[string]string{
			"RecipeID": "11",
			"FoodID":   "51",
			"Quantity": "",
			"Unit":     "ひとつまみ",
		},
		map[string]string{
			"RecipeID": "11",
			"FoodID":   "13",
			"Quantity": "",
			"Unit":     "お好みで",
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
