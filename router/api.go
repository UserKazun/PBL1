package router

import (
	"github.com/PBL1/controller"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {
	// ログイン時にクッキーにログイン情報を保持させる
	//api.POST("/auth/login", controller.PostLoginDataInCookie)

	// ログアウト時にクッキーのログイン情報を削除する
	//api.POST("/auth/logout", controller.PostLogoutDeleteCookie)

	// メニューIDを受け取り、レシピのURLを返す
	api.GET("/recipe/:menu_id", controller.GetRecipeByMenuID)
}
