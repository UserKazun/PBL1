package service

import (
	"github.com/PBL1/controller"
	"github.com/PBL1/model"
)

func GetRecipeByMenuID(menuID uint) (controller.Recipe, error) {
	modelRecipe := model.Recipe{}
	controllerRecipe := controller.Recipe{}

	err := db.Where("menu_id = ?", menuID).First(&modelRecipe).Error

	controllerRecipe.URL = modelRecipe.URL

	return controllerRecipe, err
}

