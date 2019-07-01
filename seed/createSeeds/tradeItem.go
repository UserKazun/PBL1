package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
)

// CreateSeedTradeItems ...ここで記入したデータをDBにinsertする
func CreateSeedTradeItems() {

	tradeItemsInfos := []map[string]string{
		map[string]string{
			"Name":     "スコッティ フラワーパック",
			"ImageURL": "http://ur0.work/W8jw",
			"Point":    "300",
		},
		map[string]string{
			"Name":     "クリネックス ティシュー",
			"ImageURL": "http://ur0.work/UR5I",
			"Point":    "400",
		},
		map[string]string{
			"Name":     "ウタマロ石鹸",
			"ImageURL": "http://ur0.work/02OC",
			"Point":    "450",
		},
		map[string]string{
			"Name":     "ハイディスク モバイルバッテリー5000mAh",
			"ImageURL": "http://ur0.work/Xaxx",
			"Point":    "2000",
		},
		map[string]string{
			"Name":     "プラズマクラスター搭載 卓上加湿器",
			"ImageURL": "http://ur0.work/WrVy",
			"Point":    "15000",
		},
		map[string]string{
			"Name":     "オーガニック オリーブオイル",
			"ImageURL": "http://ur0.work/HwUk",
			"Point":    "400",
		},
		map[string]string{
			"Name":     "キャノーラ油",
			"ImageURL": "http://ur0.work/3Vcm",
			"Point":    "350",
		},
		map[string]string{
			"Name":     "おーいお茶2L×6本",
			"ImageURL": "http://urx.blue/SmGu",
			"Point":    "350",
		},
		map[string]string{
			"Name":     "鼻セレブ3箱パック",
			"ImageURL": "http://urx.blue/P7lw",
			"Point":    "450",
		},
		map[string]string{
			"Name":     "電気タコ焼き機",
			"ImageURL": "http://urx.blue/VpeM",
			"Point":    "1500",
		},
		map[string]string{
			"Name":     "Dyson Cool",
			"ImageURL": "http://urx.blue/fIvJ",
			"Point":    "25000",
		},
		map[string]string{
			"Name":     "人をダメにするソファ",
			"ImageURL": "http://urx.blue/V5dk",
			"Point":    "2000",
		},
		map[string]string{
			"Name":     "SanDisc USB 128GB",
			"ImageURL": "http://urx.blue/Y0YM",
			"Point":    "1600",
		},
	}

	for _, info := range tradeItemsInfos {
		point, _ := strconv.Atoi(info["Point"])
		createTradeItem(model.TradeItem{
			Name:     info["Name"],
			ImageURL: info["ImageURL"],
			Point:    uint(point),
		})
	}
}

// createIngredient ...渡されたデータをserviceの関数へ渡し、DBにinsertさせる
func createTradeItem(tradeItem model.TradeItem) {
	tradeItem, err := service.CreateTradeItem(tradeItem)
	if err != nil {
		panic(err)
	}
	fmt.Printf("tradeItem created\n")
}
