package admin

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	V1 := r.Group("/api/v1/admin")
	{
		InitV1(V1)
	}
}
func InitV1(group *gin.RouterGroup) {

	group.GET("/tags", GetTags)
	group.POST("/tags", AddTag)
	group.PUT("/tags/:id", EditTag)
	group.DELETE("/tags/:id", DeleteTag)

	group.GET("/categories", GetTags)
	group.POST("/categories", AddTag)
	group.PUT("/categories/:id", EditTag)
	group.DELETE("/categories/:id", DeleteTag)

	group.GET("/articles", GetTags)
	group.POST("/articles", AddTag)
	group.PUT("/articles/:id", EditTag)
	group.DELETE("/articles/:id", DeleteTag)
}
