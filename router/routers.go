package router

import (
	. "github.com/blackbarth/goopportunities.git/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		InitializeHandler()
		// defer CloseDB()
		v1.GET("/opening", func(ctx *gin.Context) {
			ShowOpeningHandler(ctx)
		})
		v1.POST("/opening", func(ctx *gin.Context) {
			CreateOpeningHandler(ctx)
		})
		v1.DELETE("/opening", func(ctx *gin.Context) {
			DeleteOpeningHandler(ctx)
		})
		v1.PUT("/opening", func(ctx *gin.Context) {
			UpdateOpeningHandler(ctx)
		})
		v1.GET("/openings", func(ctx *gin.Context) {
			ListOpeningHandler(ctx)
		})
		v1.GET("/opening/:id", func(ctx *gin.Context) {
			FindOpeningHandler(ctx)
		})
	}
}
