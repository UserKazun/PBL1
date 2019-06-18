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

// PutMypageByUserID ...ユーザーIDを元にマイページ情報を更新する
func PutMypageByUserID(c *gin.Context) {
	var err error
	var currentPassword string

	userID := c.PostForm("user_id")
	userName := c.PostForm("user_name")
	userEmail := c.PostForm("user_email")
	UserStreetAddress := c.PostForm("user_street_address")
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")

	errCode := AuthCheck(c, userID)
	if errCode != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// 現在のパスワードを取得
	currentPassword, err = service.GetPasswordByID(userID)
	if err != nil {
		log.Println("データが取得できませんでした")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// パスワードを知っているか確認
	if oldPassword != currentPassword {
		log.Println("現在のパスワードと一致しません")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// ユーザ情報を更新
	err = service.PutUserByID(userID, userName, userEmail, UserStreetAddress, newPassword)
	if err != nil {
		log.Println("データが更新できませんでした")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.Println("ユーザ情報を更新しました")
	c.AbortWithStatus(http.StatusOK)
}
