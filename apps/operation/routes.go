package operation

import (
	"github.com/gin-gonic/gin"
	"zelda/apps/operation/api"
	"zelda/apps/operation/ws"
)

func InitGroup(group *gin.RouterGroup) {
	//获取标签列表
	group.GET("/tags", api.GetTags)
	group.GET("/categories", api.GetTags)
	group.GET("/articles", api.GetTags)
	group.GET("/article/:id", api.GetTags)
	group.GET("/danmu", ws.DanmuRoom)
}

func InitRoutes(r *gin.Engine) {
	V1 := r.Group("/api/v1/operation")
	{
		InitGroup(V1)
	}
}
