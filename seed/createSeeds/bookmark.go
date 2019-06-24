package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
)

// CreateSeedBookmarks ...ここで記入したデータをDBにinsertする
func CreateSeedBookmarks() {
	bookmarksInfos := []map[string]string{
		map[string]string{
			"UserID":   "goya",
			"RecipeID": "1",
		},
		map[string]string{
			"UserID":   "goya",
			"RecipeID": "4",
		},
	}

	for _, info := range bookmarksInfos {
		userID, _ := info["UserID"]
		recipeID, _ := strconv.Atoi(info["RecipeID"])
		createBookmark(model.Bookmark{
			UserID:   string(userID),
			RecipeID: uint(recipeID),
		})
	}
}

// createBookmark ...渡されたデータをserviceの関数へ渡し、DBにinsertさせる
func createBookmark(bookmark model.Bookmark) {
	bookmark, err := service.CreateBookmark(bookmark)
	if err != nil {
		panic(err)
	}
	fmt.Printf("bookmark created\n")
}
