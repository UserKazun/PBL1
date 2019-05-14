package controller

import (
	"github.com/PBL1/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetRecipes(c *gin.Context) {
	recipe, err := model.GetRecipes()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, recipe)
}
