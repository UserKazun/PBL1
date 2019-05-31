package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetRouter ...apiGroupやmiddlewareの設定
func GetRouter() *gin.Engine {
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("SessionName", store))

	api := r.Group("/api/v1")
	api2 := r.Group("/api/v2")
	apiRouter(api, api2)

	return r
}
