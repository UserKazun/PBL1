package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
)

// CreateSeedFoods ...ここで記入したデータをDBにinsertする
func CreateSeedFoods() {

	foodsInfos := []map[string]string{
		map[string]string{ //1
			"Name":  "ライス",
			"Point": "5",
		},
		map[string]string{ //2
			"Name":  "牛肉",
			"Point": "5",
		},
		map[string]string{ //3
			"Name":  "牛丼のタレ",
			"Point": "1",
		},
		map[string]string{ //4
			"Name":  "カレールー",
			"Point": "1",
		},
		map[string]string{ //5
			"Name":  "人参",
			"Point": "1",
		},
		map[string]string{ //6
			"Name":  "玉葱",
			"Point": "1",
		},
		map[string]string{ //7
			"Name":  "じゃがいも",
			"Point": "1",
		},
	}

	for _, info := range foodsInfos {
		point, _ := strconv.Atoi(info["Point"])
		createFood(model.Food{
			Name:  info["Name"],
			Point: uint(point),
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
