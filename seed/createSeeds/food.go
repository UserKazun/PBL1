package createSeeds

import (
	"fmt"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
)

// CreateSeedFoods ...ここで記入したデータをDBにinsertする
func CreateSeedFoods() {

	foodsInfos := []map[string]string{
		map[string]string{ //1
			"Name": "ライス",
		},
		map[string]string{ //2
			"Name": "牛肉",
		},
		map[string]string{ //3
			"Name": "牛丼のタレ",
		},
		map[string]string{ //4
			"Name": "カレールー",
		},
		map[string]string{ //5
			"Name": "人参",
		},
		map[string]string{ //6
			"Name": "玉葱",
		},
		map[string]string{ //7
			"Name": "じゃがいも",
		},
	}

	for _, info := range foodsInfos {
		createFood(model.Food{
			Name: info["Name"],
		})
	}
}

// createFood ...渡されたデータをserviceの関数へ渡し、DBにinsertさせる
func createFood(food model.Food) {
	food, err := service.CreateFood(food)
	if err != nil {
		panic(err)
	}
	fmt.Printf("food created\n")
}
