package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
)

// CreateSeedRecipes ...ここで記入したデータをDBにinsertする
func CreateSeedRecipes() {

	recipesInfos := []map[string]string{
		map[string]string{
			"Name":       "牛丼",
			"CategoryID": "1",
			"Price":      "150",
			"Point":      "15",
			"URL":        "hogehoge.com",
		},
		map[string]string{
			"Name":       "シチュー",
			"CategoryID": "2",
			"Price":      "200",
			"Point":      "20",
			"URL":        "fugafuga.com",
		},
		map[string]string{
			"Name":       "カレーライス",
			"CategoryID": "1",
			"Price":      "200",
			"Point":      "20",
			"URL":        "goyagoya.com",
		},
	}

	for _, info := range recipesInfos {
		categoryID, _ := strconv.Atoi(info["CategoryID"])
		price, _ := strconv.Atoi(info["Price"])
		point, _ := strconv.Atoi(info["Point"])
		createRecipe(model.Recipe{
			Name:       string(info["Name"]),
			CategoryID: uint(categoryID),
			Price:      uint(price),
			Point:      uint(point),
			URL:        string(info["URL"]),
		})
	}
}

// createRecipe ...渡されたデータをserviceの関数へ渡し、DBにinsertさせる
func createRecipe(recipe model.Recipe) {
	recipe, err := service.CreateRecipe(recipe)
	if err != nil {
		panic(err)
	}
	fmt.Printf("recipe created\n")
}
