package service

import (
	"log"

	"github.com/PBL1/model"
)

// PutCartsRecipeCount ...カート内のレシピセット数に関するテーブルを更新
func PutCartsRecipeCount(userID string, recipeID uint, recipeCount uint) error {
	beforeRecipeSetCountInCart := model.RecipeSetCountInCart{}
	var err error

	err = db.Where("user_id = ? and recipe_id = ?", userID, recipeID).First(&beforeRecipeSetCountInCart).Error
	if err != nil {
		return err
	}

	afterRecipeSetCountInCart := beforeRecipeSetCountInCart
	afterRecipeSetCountInCart.RecipeCount = &recipeCount

	err = db.Model(&beforeRecipeSetCountInCart).Update(afterRecipeSetCountInCart).Error
	if err != nil {
		return err
	}

	log.Print(afterRecipeSetCountInCart.RecipeCount)

	return nil
}

func GetRecipeSetCountInCartsByUserID(userID string) ([]model.RecipeSetCountInCart, error){
	recipeSetCountInCarts := []model.RecipeSetCountInCart{}

	err := db.Where("user_id = ?", userID).Find(&recipeSetCountInCarts).Error
	log.Println(userID)
	if err != nil {
		return nil, err
	}
	log.Printf("recipeSetCountInCarts : ", recipeSetCountInCarts)
	return recipeSetCountInCarts, nil
}

func GetRecipePriceAndPointByID(recipeID []uint) (uint, uint, error){
	recipe := model.Recipe{}
	log.Printf("recipeID :", recipeID)

	err := db.Where("id = ?", recipeID).Find(&recipe).Error
	if err != nil {
		return 0, 0, err
	}
	return recipe.Price, recipe.Point, nil
}

func InsertFoodCartContentsToPuchaseHistory(userID string, foodIDsInCarts []model.Cart, ingredientInUserCarts *model.Ingredient) error {
	foodPurchaseHistory := model.FoodPurchaseHistory{}
	ingredientInUserCart := ingredientInUserCarts
	log.Printf("foodIDsInCarts 2 :", foodIDsInCarts)
	foodIDsInCart := model.Cart{}

	// 引数で受け取ったfoodID分だけloop
	for _, foodIDsInCart = range foodIDsInCarts {
		// 該当する部分にデータを代入
		log.Printf("foodIDsInCart", foodIDsInCart)
		foodPurchaseHistory.UserID = userID
		foodPurchaseHistory.RecipeID = foodIDsInCart.RecipeID
		foodPurchaseHistory.FoodID = foodIDsInCart.FoodID
		foodPurchaseHistory.FoodCount = foodIDsInCart.FoodCount
		foodPurchaseHistory.Quantity = foodIDsInCart.FoodID * ingredientInUserCart.Quantity
		foodPurchaseHistory.Unit = ingredientInUserCart.Unit

		_, err := CreateFoodPurchaseHistory(foodPurchaseHistory)
		return err
	}
	return nil
}

func InsertRecipeCartContentsToPuchaseHistory(userID string, recipeSetCountInCarts []model.RecipeSetCountInCart , recipePrice uint, recipePoint uint) error {
	recipePurchaseHistory := model.RecipePurchaseHistory{}

	// recipeSetの分だけloop
	for _, recipeSetCountInCart := range recipeSetCountInCarts {
		// 該当する部分にデータを代入
		recipePurchaseHistory.UserID = userID
		recipePurchaseHistory.RecipeID = recipeSetCountInCart.RecipeID
		recipePurchaseHistory.RecipeCount = recipeSetCountInCart.RecipeCount
		recipePurchaseHistory.Price = recipePrice
		recipePurchaseHistory.Point = recipePoint

		_, err := CreateRecipePurchaseHistory(recipePurchaseHistory)
		return err
	}
	return nil
}