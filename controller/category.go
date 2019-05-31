package controller

import (
	"log"
	"net/http"

	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

// GetAllCategoriesName ...全てのカテゴリの名前を取得する
func GetAllCategoriesName(c *gin.Context) {
	category := Category{}
	categories := []Category{}

	modelCategories, err := service.GetAllCategoriesName()
	if err != nil {
		log.Println("カテゴリが存在しません")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, modelCategory := range modelCategories {
		category.Name = modelCategory.Name
		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, categories)
}
