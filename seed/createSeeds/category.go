package createSeeds

import (
	"fmt"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
)

// CreateSeedCategorys ...ここで記入したデータをDBにinsertする
func CreateSeedCategories() {

	categoriesInfos := []map[string]string{
		map[string]string{
			"Name": "和食",
		},
		map[string]string{
			"Name": "洋食",
		},
		map[string]string{
			"Name": "中華",
		},
		map[string]string{
			"Name": "フレンチ",
		},
		map[string]string{
			"Name": "イタリアン",
		},
		map[string]string{
			"Name": "その他",
		},
	}

	for _, info := range categoriesInfos {
		createCategory(model.Category{
			Name: info["Name"],
		})
	}
}

// createCategory ...渡されたデータをserviceの関数へ渡し、DBにinsertさせる
func createCategory(category model.Category) {
	category, err := service.CreateCategory(category)
	if err != nil {
		panic(err)
	}
	fmt.Printf("category created\n")
}
