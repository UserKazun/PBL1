package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

// GetIngredientsByRecipeID ...受け取ったレシピIDを元にそのレシピで使う材料を返す
func GetIngredientsByRecipeID(c *gin.Context) {
	ingredient := IngredientCard{}
	ingredients := IngredientCards{}

	recipeID, err := GetUint(c, "recipe_id")
	if err != nil {
		log.Println("存在しないレシピIDです")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	modelIngredients := service.GetIngredientsByRecipeID(recipeID)

	for _, modelIngredient := range modelIngredients {
		ingredient.FoodName, err = service.GetFoodNameByID(modelIngredient.FoodID)
		if err != nil {
			log.Println("存在しない食材IDです")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if modelIngredient.Quantity == 0 {
			ingredient.Quantity = modelIngredient.Unit
		} else {
			ingredient.Quantity = strconv.FormatUint(uint64(modelIngredient.Quantity), 10) + modelIngredient.Unit
		}
		ingredients = append(ingredients, ingredient)
	}

	c.JSON(http.StatusOK, ingredients)
}
