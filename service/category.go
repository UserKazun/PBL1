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

// GetAllCategoriesName ...食材IDを元に食材の名前を返す
func GetAllCategoriesName() ([]model.Category, error) {
	categories := []model.Category{}

	err := db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
