package controller

import (
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

// GetCarts ...カートの中身を取得する
func GetCarts(c *gin.Context) {
	// ModelCart ...カート
	// type Cart struct {
	// 	UserID   string `sql:"type:varchar(50)" gorm:"primary_key"`
	// 	RecipeID uint   `sql:"type:int" gorm:"primary_key"`
	// 	FoodID   uint   `sql:"type:int" gorm:"primary_key"`
	// 	Quantity uint
	// 	Unit     string
	// }
	var err error
	cart := Cart{}
	carts := []Cart{}
	food := FoodInCart{}
	foods := []FoodInCart{}
	recipeIDs := []uint{}
	modelCarts := []model.Cart{}
	modelRecipe := model.Recipe{}

	userID := c.Param("user_id")

	recipeIDs, err = service.GetRecipeIDsInCartByUserID(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	for _, recipeID := range recipeIDs {
		foods = nil
		modelRecipe = service.GetRecipeByRecipeID(recipeID)
		cart.RecipeName = modelRecipe.Name
		cart.RecipeImageURL = modelRecipe.ImageURL
		cart.Price = modelRecipe.Price
		cart.Point = modelRecipe.Point

		cart.RecipeCount, err = service.GetRecipeCount(userID, recipeID)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		modelCarts, err = service.GetFoodIDsInCartByUserID(userID, recipeID)

		for _, modelCart := range modelCarts {
			food.Name, err = service.GetFoodNameByID(modelCart.FoodID)
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			food.FoodCount = modelCart.FoodCount

			if modelCart.Quantity == 0 {
				food.Quantity = modelCart.Unit
			} else {
				food.Quantity = strconv.FormatUint(uint64(modelCart.Quantity), 10) + modelCart.Unit
			}

			foods = append(foods, food)
		}
		cart.FoodsInCart = foods

		carts = append(carts, cart)
	}

	c.JSON(http.StatusOK, carts)
}
