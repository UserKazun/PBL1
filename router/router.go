package router

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	apiRouter(api)

	return r
}
