package service

import (
	"github.com/PBL1/model"
)

// PutCartsFoodCount ...カート内の食料の数に関するテーブルを更新
func PutCartsFoodCount(userID string, recipeID uint, foodID uint, foodCount uint) error {
	BeforeCart := model.Cart{}
	var err error

	err = db.Where("user_id = ? and recipe_id = ? and food_id = ?", userID, recipeID, foodID).First(&BeforeCart).Error
	if err != nil {
		return err
	}

	afterCart := BeforeCart
	afterCart.FoodCount = &foodCount

	err = db.Model(&BeforeCart).Update(afterCart).Error
	if err != nil {
		return err
	}

	return nil
}
