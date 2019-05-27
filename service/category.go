package service

import (
	"github.com/PBL1/model"
)

// CreateCategory ...DBに与えられたデータをinsertする
func CreateCategory(category model.Category) (model.Category, error) {
	err := db.Create(&category).Error
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}
