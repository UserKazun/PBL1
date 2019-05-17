package controller

import (
	"log"
	"net/http"

	"github.com/PBL1/service"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var session sessions.Session //API別にセッションがリセットされるためグローバルで保存

// PostLoginDataInCookie ...ログイン時にセッションにログイン情報を保持させる
func PostLoginDataInCookie(c *gin.Context) {

	loginUser := LoginUser{}

	userID := c.PostForm("user_id")
	password := c.PostForm("password")

	user, err := service.CheckAccount(userID, password)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	service.PostTrueToAdmin(userID)

	loginUser.ID = user.ID
	loginUser.Name = user.Name
	loginUser.IsAdmin = true

	session = sessions.Default(c)

	//ログインIDをKeyにセッションIDを生成
	session.Set(userID, userID)

	session.Save()

	c.JSON(http.StatusOK, loginUser)

}

// PostLogoutDeleteCookie ...ログアウト時にセッションからログイン情報を削除する
func PostLogoutDeleteCookie(c *gin.Context) {

	userID := c.PostForm("user_id")

	sessionUser := session.Get(userID)

	if sessionUser == nil {
		log.Println("セッションがありません")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	session.Delete(userID)
	session.Save()

	c.AbortWithStatus(http.StatusOK)
}
