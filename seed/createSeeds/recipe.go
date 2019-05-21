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
			"ID":       "1",
			"Name":     "牛丼",
			"Category": "1",
			"URL":      "hogehoge.com",
		},
		map[string]string{
			"ID":       "2",
			"Name":     "シチュー",
			"Category": "2",
			"URL":      "fugafuga.com",
		},
		map[string]string{
			"ID":       "3",
			"Name":     "カレーライス",
			"Category": "1",
			"URL":      "goyagoya.com",
		},
	}

	for _, info := range recipesInfos {
		ID, _ := strconv.Atoi(info["ID"])
		Name, _ := info["Name"]
		category, _ := strconv.Atoi(info["Category"])
		URL, _ := info["URL"]
		createRecipe(model.Recipe{
			ID:       uint(ID),
			Name:     string(Name),
			Category: uint(category),
			URL:      string(URL),
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
