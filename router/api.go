package router

import (
	"github.com/PBL1/controller"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup, api2 *gin.RouterGroup) {

	// ログイン時にセッションとしてログイン情報を保持させる
	api.POST("/auth/login", controller.PostLoginDataInCookie)

	// ログアウト時にセッションとしてログイン情報を削除する
	api.POST("/auth/logout", controller.PostLogoutDeleteCookie)

	// 受け取ったrecipeIDを元にレシピのURLを取得する
	api.GET("/recipes/:recipe_id/URL", controller.GetRecipeByMenuID)

	// 受け取ったレシピIDを元にそのレシピ画面を表示するのに必要なデータ（材料）を取得する
	api.GET("/ingredient/:recipe_id", controller.GetIngredientsByRecipeID)

	// 受け取ったレシピIDを元にそのレシピ画面を表示するのに必要なデータ（材料・画像URL）を取得する
	api2.GET("/recipes/:recipe_id", controller.GetRecipeByRecipeID)

	// 受け取った材料IDを指定されたユーザーのカートに追加する
	api.POST("/carts", controller.PostIngredientsToCart)

	// 全てのカテゴリの名前を取得する
	api.GET("/categories", controller.GetAllCategoriesName)

	// 与えられたキーを元に検索した結果のレシピデータを取得する
	api.GET("search-recipes/categories/:category_id/keys/:search_key", controller.GetRecipesSearch)
}
