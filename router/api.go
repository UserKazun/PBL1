package router

import (
	"github.com/PBL1/controller"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {

	// ログイン時にセッションとしてログイン情報を保持させる
	api.POST("/auth/login", controller.PostLoginDataInCookie)

	// ログアウト時にセッションとしてログイン情報を削除する
	api.POST("/auth/logout", controller.PostLogoutDeleteCookie)

	// 受け取ったrecipeIDを元にレシピのURLを取得する
	api.GET("/recipe/:recipe_id/URL", controller.GetRecipeByMenuID)

	// 受け取ったレシピIDを元にそのレシピに必要な材料を取得する
	api.GET("/ingredient/:recipe_id", controller.GetIngredientsByRecipeID)

	// 受け取った材料IDを指定されたユーザーのカートに追加する
	api.POST("/cart", controller.PostIngredientsToCart)

	// 全てのカテゴリの名前を取得する
	api.GET("/categories", controller.GetAllCategoriesName)
}
