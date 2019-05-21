package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostIngredientsToCart ...受け取った材料IDを指定されたユーザーのカートに追加する
func PostIngredientsToCart(c *gin.Context) {
	userID := c.PostForm("user_id")

	err := AuthCheck(c, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.AbortWithStatus(http.StatusOK)
	//c.JSON(http.StatusOK)
}
