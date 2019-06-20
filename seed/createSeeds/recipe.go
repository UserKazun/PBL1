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
			"Name":        "牛丼",
			"Description": "お肉がのったご飯だよ！",
			"CategoryID":  "2",
			"ImageURL":    "https://www.pakutaso.com/shared/img/thumb/KAZUHIRO171013022_TP_V.jpg",
			"PageURL":     "hogehoge.com",
			"Price":       "150",
			"Point":       "15",
		},
		map[string]string{
			"Name":        "肉だけカレーライス",
			"Description": "男は黙って肉食っとけカレーライス",
			"CategoryID":  "7",
			"ImageURL":    "https://d2l930y2yx77uc.cloudfront.net/production/uploads/images/7120502/picture_pc_bd3805fab5e332c67b1862c988179471.jpg",
			"PageURL":     "fugafuga.com",
			"Price":       "200",
			"Point":       "20",
		},
		map[string]string{
			"Name":        "シチュー",
			"Description": "",
			"CategoryID":  "5",
			"ImageURL":    "https://d2dcan0armyq93.cloudfront.net/photo/odai/600/222569875b57db9b87ae55845b35315d_600.jpg",
			"PageURL":     "fugafuga.com",
			"Price":       "200",
			"Point":       "20",
		},
		map[string]string{
			"Name":        "米食っとけ、そう米だけさ",
			"Description": "米オンリー",
			"CategoryID":  "7",
			"ImageURL":    "https://d2dcan0armyq93.cloudfront.net/photo/odai/600/222569875b57db9b87ae55845b35315d_600.jpg",
			"PageURL":     "komekome.com",
			"Price":       "50",
			"Point":       "5",
		},
	}

	for _, info := range recipesInfos {
		categoryID, _ := strconv.Atoi(info["CategoryID"])
		price, _ := strconv.Atoi(info["Price"])
		point, _ := strconv.Atoi(info["Point"])
		createRecipe(model.Recipe{
			Name:        info["Name"],
			CategoryID:  uint(categoryID),
			ImageURL:    info["ImageURL"],
			PageURL:     info["PageURL"],
			Price:       uint(price),
			Point:       uint(point),
			Description: info["Description"],
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
