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
