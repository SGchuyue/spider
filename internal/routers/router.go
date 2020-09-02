package routers
import (
	"github.com/gin-gonic/gin"
//	"writetest/internal/routers/api"
	"writetest/internal/routers/api/v1"

)
func NewRouter() *gin.Engine {
	r := gin.New()
	article := v1.NewArticle()
	title := v1.NewTitle()

	apiv1 := r.Group("/api/v1")
	//apiv1.Use() //middleware.JWT()
	{
		// 创建标题
		apiv1.POST("/titles", title.Create)
		// 删除指定标题
		apiv1.DELETE("/titles/:id", title.Delete)
		// 更新标题标签
		apiv1.PUT("/titles/:id", title.Update)

		// 创建漏洞
		apiv1.POST("/articles", article.Create)
		// 删除指定漏洞
		apiv1.DELETE("/articles/:id", article.Delete)
		// 更新指定漏洞
		apiv1.PUT("/articles/:id", article.Update)
		// 获取指定漏洞
//		apiv1.GET("/articles/:id", article.Get)

	}

	return r
}
