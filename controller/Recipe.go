package controller

import (
	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetRecipeByMenuID(c *gin.Context) {
	var menuID uint
	var err error

	recipe := Recipe{}

	menuID, err = GetUint(c,"menu_id")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

	recipe, err = service.GetRecipeByMenuID(menuID)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, recipe)
}
