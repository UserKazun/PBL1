package controller

import (
	"log"
	"net/http"

	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

// GetMypageByUserID ...ユーザーIDを元にマイページ情報を取得する
func GetMypageByUserID(c *gin.Context) {
	mypage := Mypage{}

	userID := c.Param("user_id")

	errCode := AuthCheck(c, userID)
	if errCode != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := service.GetUserByID(userID)
	if err != nil {
		log.Println("データが取得できませんでした")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	mypage.UserID = user.ID
	mypage.UserName = user.Name
	mypage.UserEmail = user.Email
	mypage.UserStreetAddress = user.StreetAddress
	mypage.CumulativePoints, err = service.GetCumulativePointsByUserID(userID)
	if err != nil {
		log.Println("データが取得できませんでした")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, mypage)
}
