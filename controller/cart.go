package controller

import (
	"net/http"
	"strconv"

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

	c.AbortWithStatus(http.StatusOK)
}
