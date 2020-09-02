package routers

import (
	"github.com/gin-gonic/gin"
	//	"myproject.cn/spider/internal/routers/api"
	"myproject.cn/spider/http_server/routers/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	article := api.NewArticle()
	title := api.NewTitle()

	router := r.Group("/api")
	//router.Use() //middleware.JWT()
	{
		// 创建标题
		router.POST("/titles", title.Create)
		// 删除指定标题
		router.DELETE("/titles/:id", title.Delete)
		// 更新标题标签
		router.PUT("/titles/:id", title.Update)

		// 创建漏洞
		router.POST("/articles", article.Create)
		// 删除指定漏洞
		router.DELETE("/articles/:id", article.Delete)
		// 更新指定漏洞
		router.PUT("/articles/:id", article.Update)
		// 获取指定漏洞
		//		router.GET("/articles/:id", article.Get)
	}

	return r
}
