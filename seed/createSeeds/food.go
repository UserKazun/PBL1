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
		map[string]string{ //8
			"Name": "卵",
		},
		map[string]string{ //9
			"Name": "パスタ",
		},
		map[string]string{ //10
			"Name": "ベーコン",
		},
		map[string]string{ //11
			"Name": "ニンニク",
		},
		map[string]string{ //12
			"Name": "粉チーズ",
		},
		map[string]string{ //13
			"Name": "黒コショウ",
		},
		map[string]string{ //14
			"Name": "ピュアオイル",
		},
		map[string]string{ //15
			"Name": "豆腐",
		},
		map[string]string{ //16
			"Name": "ひき肉",
		},
		map[string]string{ //17
			"Name": "醤油",
		},
		map[string]string{ //18
			"Name": "豆板醤",
		},
		map[string]string{ //19
			"Name": "塩コショウ",
		},
		map[string]string{ //20
			"Name": "片栗粉",
		},
		map[string]string{ //21
			"Name": "エビ",
		},
		map[string]string{ //22
			"Name": "ブロッコリー",
		},
		map[string]string{ //23
			"Name": "スイートコーン",
		},
		map[string]string{ //24
			"Name": "すりおろしにんにく",
		},
		map[string]string{ //25
			"Name": "オリーブオイル",
		},
		map[string]string{ //26
			"Name": "とけるチーズ",
		},
		map[string]string{ //27
			"Name": "タルタルソース",
		},
		map[string]string{ //28
			"Name": "パン粉",
		},
		map[string]string{ //29
			"Name": "甘塩鮭",
		},
		map[string]string{ //30
			"Name": "しめじ",
		},
		map[string]string{ //31
			"Name": "料理酒",
		},
		map[string]string{ //32
			"Name": "バター",
		},
		map[string]string{ //33
			"Name": "青じそドレッシング",
		},
		map[string]string{ //34
			"Name": "きゅうり",
		},
		map[string]string{ //35
			"Name": "レタス",
		},
		map[string]string{ //36
			"Name": "ワカメ",
		},
		map[string]string{ //37
			"Name": "トマト",
		},
		map[string]string{ //38
			"Name": "砂糖",
		},
		map[string]string{ //39
			"Name": "酢",
		},
		map[string]string{ //40
			"Name": "鶏ガラスープの素",
		},
		map[string]string{ //41
			"Name": "ごま油",
		},
		map[string]string{ //42
			"Name": "にんにくチューブ",
		},
		map[string]string{ //43
			"Name": "いりごま",
		},
		map[string]string{ //44
			"Name": "鶏もも肉",
		},
		map[string]string{ //45
			"Name": "みりん",
		},
		map[string]string{ //46
			"Name": "味の素",
		},
		map[string]string{ //47
			"Name": "粉山椒",
		},
		map[string]string{ //48
			"Name": "刻みネギ",
		},
		map[string]string{ //49
			"Name": "豆苗",
		},
		map[string]string{ //50
			"Name": "めんつゆ",
		},
		map[string]string{ //51
			"Name": "塩",
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
