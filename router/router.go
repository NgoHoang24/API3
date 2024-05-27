package router

import (
	"API3/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter(tagshandler *handler.TagsHandler) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagshandler.FindAll)
	tagsRouter.GET("/:tagId", tagshandler.FindByID)
	tagsRouter.POST("", tagshandler.Create)
	tagsRouter.PATCH("/:tagId", tagshandler.Update)
	tagsRouter.DELETE("/:tagId", tagshandler.Delete)

	return router
}
