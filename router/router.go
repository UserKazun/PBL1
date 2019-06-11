package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/gin-contrib/cors"
)

// GetRouter ...apiGroupやmiddlewareの設定
func GetRouter() *gin.Engine {
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("SessionName", store))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	api := r.Group("/api/v1")
	apiRouter(api)

	return r
}
