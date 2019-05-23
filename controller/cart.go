package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

// PostIngredientsToCart ...受け取った材料IDを指定されたユーザーのカートに追加する
func PostIngredientsToCart(c *gin.Context) {

	// Cart ...カート
	// type Cart struct {
	// 	UserID   string `sql:"type:varchar(50)" gorm:"primary_key"`
	// 	RecipeID uint   `sql:"type:int" gorm:"primary_key"`
	// 	FoodID   uint
	// 	Quantity uint
	// 	Unit     string
	// }

	// Ingredient ...材料
	// type Ingredient struct {
	// 	RecipeID uint `sql:"type:int" gorm:"primary_key"`
	// 	FoodID   uint `sql:"type:int" gorm:"primary_key"`
	// 	Quantity uint
	// 	Unit     string
	// }

	userID := c.PostForm("user_id")
	recipeID, _ := strconv.Atoi(c.PostForm("recipe_id"))
	uintRecipeID := uint(recipeID)

	errCode := AuthCheck(c, userID)
	if errCode != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	modelIngredients, err := service.GetIngredientsByRecipeID(uintRecipeID)
	if err != nil {
		log.Println("存在しないレシピIDです")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = service.PostIngredientsToCart(userID, modelIngredients)
	if err != nil {
		log.Println("与えられたパラメータが適切ではありません")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.AbortWithStatus(http.StatusOK)

}
