package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

// PostIngredientsToCart ...受け取った材料IDを指定されたユーザーのカートに追加する
func PostIngredientsToCart(c *gin.Context) {
	userID := c.PostForm("user_id")
	recipeID, _ := strconv.Atoi(c.PostForm("recipe_id"))
	uintRecipeID := uint(recipeID)

	errCode := AuthCheck(c, userID)
	if errCode != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	modelIngredients := service.GetIngredientsByRecipeID(uintRecipeID)

	service.PostIngredientsToCart(userID, modelIngredients)
	service.InsertRecipeCount(userID, uintRecipeID)

	c.AbortWithStatus(http.StatusOK)
}

// GetCartsByUserID ...カートの中身を取得する
func GetCartsByUserID(c *gin.Context) {
	var err error
	cart := Cart{}
	carts := []Cart{}
	food := Food{}
	foods := []Food{}
	recipeIDs := []uint{}
	modelCarts := []model.Cart{}
	modelRecipe := model.Recipe{}
	ingredient := &model.Ingredient{}

	userID := c.Param("user_id")

	errCode := AuthCheck(c, userID)
	if errCode != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	recipeIDs, err = service.GetRecipeIDsInCartByUserID(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, recipeID := range recipeIDs {
		foods = nil
		cart.RecipeID = recipeID
		modelRecipe = service.GetRecipeByRecipeID(recipeID)
		cart.RecipeName = modelRecipe.Name
		cart.RecipeImageURL = modelRecipe.ImageURL
		cart.Price = modelRecipe.Price
		cart.Point = modelRecipe.Point

		cart.RecipeCount, err = service.GetRecipeCount(userID, recipeID)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		modelCarts, err = service.GetFoodIDsInCartByUserID(userID, recipeID)

		for _, modelCart := range modelCarts {
			food.ID = modelCart.FoodID
			food.Name, err = service.GetFoodNameByID(modelCart.FoodID)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			food.FoodCount = *modelCart.FoodCount

			ingredient, err = service.GetIngredientsByRecipeIDAndFoodID(recipeID, modelCart.FoodID)

			if ingredient.Quantity == 0 {
				food.Quantity = ingredient.Unit
			} else {
				food.Quantity = strconv.FormatUint(uint64(ingredient.Quantity), 10) + ingredient.Unit
			}

			foods = append(foods, food)
		}
		cart.Foods = foods

		carts = append(carts, cart)
	}

	c.JSON(http.StatusOK, carts)
}

// PutCartsRecipeCountByUserID ...カートの中身(レシピの数)を更新する
func PutCartsRecipeCountByUserID(c *gin.Context) {
	userID := c.PostForm("user_id")
	recipeID, _ := strconv.Atoi(c.PostForm("recipe_id"))
	recipeCount, _ := strconv.Atoi(c.PostForm("recipe_count"))

	errCode := AuthCheck(c, userID)
	if errCode != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err := service.PutCartsRecipeCount(userID, uint(recipeID), uint(recipeCount))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

// PutCartsFoodCountByUserID ...カートの中身(食料の数)を更新する
func PutCartsFoodCountByUserID(c *gin.Context) {
	userID := c.PostForm("user_id")
	recipeID, _ := strconv.Atoi(c.PostForm("recipe_id"))
	foodID, _ := strconv.Atoi(c.PostForm("food_id"))
	foodCount, _ := strconv.Atoi(c.PostForm("food_count"))

	errCode := AuthCheck(c, userID)
	if errCode != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err := service.PutCartsFoodCount(userID, uint(recipeID), uint(foodID), uint(foodCount))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)

}

func DeleteCartContentByUserIDandRecipeID(c *gin.Context) {
	userID := c.Param("user_id")
	recipeID, err := GetUint(c, "recipe_id")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	errCode := AuthCheck(c, userID)
	if errCode != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = service.DeleteCartContentByUserIDandRecipeID(userID, recipeID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
