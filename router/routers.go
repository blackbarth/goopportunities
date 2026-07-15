package router

import (
	docs "github.com/blackbarth/goopportunities.git/docs"
	. "github.com/blackbarth/goopportunities.git/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRouter(router *gin.Engine) {
	InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{

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

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json")))
	// docs.SwaggerInfo.BasePath = "/api/v1"
}
