package controller

import (
	"log"
	"net/http"

	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

// GetRecipeByMenuID ...受け取ったメニューIDを元にレシピのURLを返す
func GetRecipeByMenuID(c *gin.Context) {
	var recipeID uint
	var err error
	var recipeURL string

	recipeID, err = GetUint(c, "recipe_id")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	recipeURL, err = service.GetRecipeByMenuID(recipeID)
	if err != nil {
		log.Println("そのレシピは存在しません")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"recipe_URL": recipeURL,
	})
}
