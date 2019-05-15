package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/PBL1/model"
)

func CreateSeedRecipes() {

	recipes_infos := []map[string]string{
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

	for _, info := range recipes_infos {
		menuID, _ := strconv.Atoi(info["menuID"])
		URL, _ := strconv.Atoi(info["URL"])
		createResipe(model.Recipe{
			MenuID: uint(menuID),
			URL:    string(URL),
		})
	}
}

func createResipe(recipe model.Recipe) {
	recipe, err := service.CreateRecipe(recipe)
	if err != nil {
		panic(err)
	}
	fmt.Printf("recipe created\n")
}
