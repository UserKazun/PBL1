package model


func CreateRecipe(recipe Recipe) (Recipe, error) {
	err := db.Create(&recipe).Error
	if err != nil {
		return Recipe{}, err
	}
	return recipe, nil
}