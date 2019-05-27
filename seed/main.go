package main

import "github.com/PBL1/seed/createSeeds"

func main() {
	createSeeds.CreateSeedRecipes()
	createSeeds.CreateSeedUsers()
	createSeeds.CreateSeedFoods()
	createSeeds.CreateSeedIngredients()
	createSeeds.CreateSeedCategories()

}
