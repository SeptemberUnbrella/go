package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-programming-tour-book/blog-service/internal/api/v1"
)

func NewRouter() *gin.Engine {
	//r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	//apiv1 := r.Group("/api/v1")
	//{
	//	apiv1.POST("/tags")
	//	apiv1.DELETE("/tags/:id")
	//	apiv1.PUT("/tags/:id")
	//	apiv1.PATCH("/tags/:id/state")
	//	apiv1.GET("/tags")
	//
	//	apiv1.POST("/articles")
	//	apiv1.DELETE("/articles/:id")
	//	apiv1.PUT("/articles/:id")
	//	apiv1.PATCH("/articles/:id/state")
	//	apiv1.GET("/articles")
	//
	//}
	//return r

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
