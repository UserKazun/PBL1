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
		map[string]string{ //1
			"Name":        "牛丼",
			"Description": "お肉がのったご飯だよ！",
			"CategoryID":  "2",
			"ImageURL":    "https://www.pakutaso.com/shared/img/thumb/KAZUHIRO171013022_TP_V.jpg",
			"PageURL":     "hogehoge.com",
			"Price":       "150",
			"Point":       "15",
		},
		map[string]string{ //2
			"Name":        "肉だけカレーライス",
			"Description": "男は黙って肉食っとけカレーライス",
			"CategoryID":  "7",
			"ImageURL":    "https://d2l930y2yx77uc.cloudfront.net/production/uploads/images/7120502/picture_pc_bd3805fab5e332c67b1862c988179471.jpg",
			"PageURL":     "fugafuga.com",
			"Price":       "200",
			"Point":       "20",
		},
		map[string]string{ //3
			"Name":        "シチュー",
			"Description": "",
			"CategoryID":  "5",
			"ImageURL":    "https://d2dcan0armyq93.cloudfront.net/photo/odai/600/222569875b57db9b87ae55845b35315d_600.jpg",
			"PageURL":     "fugafuga.com",
			"Price":       "200",
			"Point":       "20",
		},
		map[string]string{ //4
			"Name":        "米食っとけ、そう米だけさ",
			"Description": "米オンリー",
			"CategoryID":  "7",
			"ImageURL":    "https://d2dcan0armyq93.cloudfront.net/photo/odai/600/222569875b57db9b87ae55845b35315d_600.jpg",
			"PageURL":     "komekome.com",
			"Price":       "50",
			"Point":       "5",
		},

		map[string]string{ //5
			"Name":        "カルボナーラ",
			"Description": "生クリームを使わない本格カルボナーラ",
			"CategoryID":  "6",
			"ImageURL":    "https://img.cpcdn.com/recipes/5690951/m/89fb15fa9fc5a2f2ab69f7119cabdc72.jpg?u=31342300&p=1560167364",
			"Price":       "200",
			"Point":       "20",
		},
		map[string]string{ //6
			"Name":        "麻婆豆腐",
			"Description": "ご飯がススム麻婆豆腐",
			"CategoryID":  "4",
			"ImageURL":    "https://img.cpcdn.com/recipes/5703427/m/042ace326144a23e557155b9b8b7626f.jpg?u=27509585&p=1560855608",
			"Price":       "180",
			"Point":       "18",
		},
		map[string]string{ //7
			"Name":        "洋風エビとブロッコリーのタルタルグラタン",
			"Description": "簡単レシピでおかずにもなる！",
			"CategoryID":  "3",
			"ImageURL":    "https://img.cpcdn.com/recipes/5203054/m/1768fd5e0ccb5ff0e99d7e1ad28cc4c7.jpg?u=5656434&p=1533897825",
			"Price":       "250",
			"Point":       "25",
		},
		map[string]string{ //8
			"Name":        "塩鮭とたまねぎのバター蒸し",
			"Description": "電子レンジで簡単、玉ねぎの甘みが美味しい",
			"CategoryID":  "7",
			"ImageURL":    "https://img.cpcdn.com/recipes/1338962/m/582bed2d69e204023c6e1ba624e242e0.jpg?u=1559749&p=1295872906",
			"Price":       "300",
			"Point":       "30",
		},
		map[string]string{ //9
			"Name":        "豆腐の韓国風",
			"Description": "電子レンジで簡単、玉ねぎの甘みが美味しい",
			"CategoryID":  "7",
			"ImageURL":    "https://img.cpcdn.com/recipes/4491418/m/deedaf2ac50639cae185d7792e74869e.jpg?u=11000068&p=1501231688",
			"Price":       "150",
			"Point":       "15",
		},
		map[string]string{ //10
			"Name":        "親子丼",
			"Description": "一人前130円で簡単",
			"CategoryID":  "2",
			"ImageURL":    "https://img.cpcdn.com/recipes/5583379/m/4202a153a1a4856ddd6deb670bf775d9.jpg?u=15874183&p=1557278663",
			"Price":       "130",
			"Point":       "13",
		},
		map[string]string{ //11
			"Name":        "炒飯",
			"Description": "一人暮らしの簡単炒飯",
			"CategoryID":  "4",
			"ImageURL":    "https://img.cpcdn.com/recipes/5491104/m/ec26cd2406f1236d24ac65a1e5284c7e.jpg?u=11613384&p=1549431395",
			"Price":       "100",
			"Point":       "10",
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
