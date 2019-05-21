package controller

import (
	"log"
	"net/http"

	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

// GetRecipeByMenuID ...受け取ったメニューIDを元にレシピのURLを返す
func GetRecipeByMenuID(c *gin.Context) {
	var menuID uint
	var err error

	recipe := Recipe{}

	menuID, err = GetUint(c, "menu_id")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	recipe.URL, err = service.GetRecipeByMenuID(menuID)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, recipe)
}
