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
			"menuID": "1",
			"URL":    "hogehoge.com",
		},
		map[string]string{
			"menuID": "2",
			"URL":    "fugafuga.com",
		},
		map[string]string{
			"menuID": "3",
			"URL":    "goyagoya.com",
		},
	}

	for _, info := range recipesInfos {
		menuID, _ := strconv.Atoi(info["menuID"])
		URL, _ := info["URL"]
		createRecipe(model.Recipe{
			MenuID: uint(menuID),
			URL:    string(URL),
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
