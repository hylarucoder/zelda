package user

import (
	"github.com/gin-gonic/gin"
	"zelda/apps/user/api"
)

func InitRoutes(r *gin.Engine) {
	V1 := r.Group("/api/v1/user")
	{
		InitV1(V1)
	}
}
func InitV1(group *gin.RouterGroup) {

	group.GET("/tags", api.GetTags)
	group.POST("/tags", api.AddTag)
	group.PUT("/tags/:id", api.EditTag)
	group.DELETE("/tags/:id", api.DeleteTag)

	group.GET("/categories", api.GetTags)
	group.POST("/categories", api.AddTag)
	group.PUT("/categories/:id", api.EditTag)
	group.DELETE("/categories/:id", api.DeleteTag)

	group.GET("/articles", api.GetTags)
	group.POST("/articles", api.AddTag)
	group.PUT("/articles/:id", api.EditTag)
	group.DELETE("/articles/:id", api.DeleteTag)
}
