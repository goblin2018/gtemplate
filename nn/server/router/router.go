package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})
	return r
}
