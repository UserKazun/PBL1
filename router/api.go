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

	// キーを元に検索した結果のレシピデータを取得する(カテゴリー検索のみ)
	api.GET("search-recipes/categories/:category_id", controller.GetRecipesCategoryOnlySearch)

	// 対象ユーザのカートの中身を取得する
	api.GET("/carts/users/:user_id", controller.GetCartsByUserID)

	// 対象ユーザの中身(レシピ)を更新する
	api.PUT("/carts/recipe-counts", controller.PutCartsRecipeCountByUserID)

	// 対象ユーザの中身(食材)を更新する
	api.PUT("/carts/food-counts", controller.PutCartsFoodCountByUserID)

	// 対象ユーザのカートの中身を購入する
	api.POST("/purchase-histories", controller.PostPurchaseHistoriesByUserID)

	// 対象ユーザーのカートの中身を削除する
	api.POST("/delete-cart-content", controller.PostDeleteCartContentByUserID)

	// 対象ユーザの購入履歴を取得する
	api.GET("/purchase-histories/users/:user_id", controller.GetPurchaseHistoriesByUserID)

	// 対象ユーザのマイページ情報を取得する
	api.POST("/mypage", controller.GetMypageByUserID)

	// 対象ユーザのマイページ情報を更新する
	api.PUT("/mypage", controller.PutMypageByUserID)

	// 対象ユーザのブックマーク情報を取得する
	api.GET("/bookmark/users/:user_id", controller.GetBookmarkByUserID)

	// 対象ユーザのブックマーク情報を追加する
	api.POST("/bookmark", controller.PostBookmarkByUserID)

	// 対象ユーザのブックマーク情報を削除する
	api.DELETE("/bookmark/users/:user_id/recipes/:recipe_id", controller.DeleteBookmarkByUserID)

	// 飢餓貢献ポイントで交換できる商品情報を取得する
	api.GET("/trade-items", controller.GetTradeItems)

	// 対象ユーザが飢餓貢献ポイントで商品を交換する
	api.POST("/trade-items", controller.PostTradeItemsByUserID)

	// 飢餓貢献ポイントで交換した履歴情報を取得する
	api.GET("/trade-items-Histories/users/:user_id", controller.GetTradeItemsHistoriesByUserID)
}
