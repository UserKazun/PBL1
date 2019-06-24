package service

import (
	"github.com/PBL1/model"
)

// CreateBookmark ...DBに与えられたデータをinsertする
func CreateBookmark(bookmark model.Bookmark) (model.Bookmark, error) {
	err := db.Create(&bookmark).Error
	if err != nil {
		return model.Bookmark{}, err
	}
	return bookmark, nil
}

// GetBookmarkRecipeIDsByUserID ...
func GetBookmarkRecipeIDsByUserID(userID string) ([]uint, error) {
	bookmarks := []model.Bookmark{}
	var recipeIDs []uint

	err := db.Where("user_id = ?", userID).Find(&bookmarks).Error
	if err != nil {
		return nil, err
	}

	for _, bookmark := range bookmarks {
		recipeIDs = append(recipeIDs, bookmark.RecipeID)
	}

	return recipeIDs, nil
}
