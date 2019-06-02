package user

import "github.com/gin-gonic/gin"

func InitGroup(group *gin.RouterGroup) {
	//获取标签列表
	group.GET("/tags", GetTags)
	group.GET("/categories", GetTags)
	group.GET("/articles", GetTags)
	group.GET("/article/:id", GetTags)
	group.GET("/danmu", DanmuRoom)
}

func InitRoutes(r *gin.Engine) {
	V1 := r.Group("/api/v1/user")
	{
		InitGroup(V1)
	}
}
