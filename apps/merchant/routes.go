package merchant

import (
	"github.com/gin-gonic/gin"
	"zelda/apps/merchant/api"
)

func InitGroup(group *gin.RouterGroup) {
	//获取标签列表
	group.GET("/tags", api.GetTags)
	group.GET("/categories", api.GetTags)
	group.GET("/articles", api.GetTags)
	group.GET("/article/:id", api.GetTags)
}

func InitRoutes(r *gin.Engine) {
	V1 := r.Group("/api/v1/merchant")
	{
		InitGroup(V1)
	}
}
