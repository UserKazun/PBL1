package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetRouter ...apiGroupやmiddlewareの設定
func GetRouter() *gin.Engine {
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("SessionName", store))

	r.Use(cors.Default())

	api := r.Group("/api/v1")
	apiRouter(api)

	return r
}
