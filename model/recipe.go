package model

import "github.com/PBL1/service"

/*
フロント側からrecipe IDが送られてくるので、それに値するURLをDBから返す
*/

func GetRecipes() (service.Recipe, error) {
	recipe := service.Recipe{}
	modelrecipe := Recipe{}

	recipe.MenuID = modelrecipe.MenuID

	err := db.First(&recipe).Error

	return recipe, err
}
