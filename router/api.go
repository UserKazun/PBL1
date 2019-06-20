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

	// recipeIDを元にレシピのURLを取得する
	api.GET("/recipes/:recipe_id/URL", controller.GetRecipeByMenuID)

	// レシピIDを元にそのレシピ画面を表示するのに必要なデータ（材料）を取得する
	api.GET("/ingredient/:recipe_id", controller.GetIngredientsByRecipeID)

	// レシピIDを元にそのレシピ画面を表示するのに必要なデータ（材料・画像URL）を取得する
	api2.GET("/recipes/:recipe_id", controller.GetRecipeByRecipeID)

	// 材料を指定されたユーザーのカートに追加する
	api.POST("/carts", controller.PostIngredientsToCart)

	// 全てのカテゴリの名前を取得する
	api.GET("/categories", controller.GetAllCategoriesName)

	// キーを元に検索した結果のレシピデータを取得する
	api.GET("search-recipes/categories/:category_id/keys/:search_key", controller.GetRecipesSearch)

	// 対象ユーザのカートの中身を取得する
	api.GET("/carts/users/:user_id", controller.GetCartsByUserID)

	// 対象ユーザの中身(レシピ)を更新する
	api.PUT("/carts/recipe-counts", controller.PutCartsRecipeCountByUserID)

	// 対象ユーザの中身(食材)を更新する
	api.PUT("/carts/food-counts", controller.PutCartsFoodCountByUserID)

	// 対象ユーザのカートの中身を購入する
	//api.POST("/purchase-histories", controller.PostPurchaseHistoriesByUserID)

	// 対象ユーザの購入履歴を取得する
	api.GET("/purchase-histories/users/:user_id", controller.GetPurchaseHistoriesByUserID)

	// 対象ユーザのマイページ情報を取得する
	api.POST("/mypage", controller.GetMypageByUserID)

	// 対象ユーザのマイページ情報を更新する
	api.PUT("/mypage", controller.PutMypageByUserID)

	// 対象ユーザのブックマーク情報を取得する
	api.GET("/bookmark/:user_id", controller.GetBookmarkByUserID)

}
