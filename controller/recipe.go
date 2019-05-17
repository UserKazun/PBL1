package controller

import (
	"log"
	"net/http"

	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

// GetRecipeByMenuID ...メニューIDを受け取り、レシピのURLを返す
func GetRecipeByMenuID(c *gin.Context) {
	var menuID uint
	var err error

	recipe := Recipe{}

	menuID, err = GetUint(c, "menu_id")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

	recipe.URL, err = service.GetRecipeByMenuID(menuID)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, recipe)
}
