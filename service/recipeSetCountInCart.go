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

func GetRecipeSetCountInCartsByUserID(userID string) (*model.RecipeSetCountInCart, error) {
	recipeSetCountInCart := model.RecipeSetCountInCart{}

	err := db.Where("user_id = ?", userID).Find(&recipeSetCountInCart).Error
	log.Println(userID)
	if err != nil {
		return nil, err
	}

	return &recipeSetCountInCart, nil
}

func InsertRecipeCartContentsToPuchaseHistory(userID string, recipeID uint, recipeSetCountInCart model.RecipeSetCountInCart, recipePrice uint, recipePoint uint) (uint, error) {
	recipePurchaseHistory := model.RecipePurchaseHistory{}
	var err error

	// 該当する部分にデータを代入
	recipePurchaseHistory.UserID = userID
	recipePurchaseHistory.RecipeID = recipeID
	recipePurchaseHistory.RecipeCount = recipeSetCountInCart.RecipeCount
	recipePurchaseHistory.Price = recipePrice
	recipePurchaseHistory.Point = recipePoint

	_, err = CreateRecipePurchaseHistory(recipePurchaseHistory)
	if err != nil {
		return 0, err
	}

	err = db.Raw("select * from recipe_purchase_histories order by id").Scan(&recipePurchaseHistory).Error
	if err != nil {
		return 0, err
	}
	log.Println("recipePurchaseHistory.ID", recipePurchaseHistory.ID)

	return recipePurchaseHistory.ID, nil
}
