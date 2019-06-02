package misc

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	r.GET("/", Home)
	r.GET("/raw", NowPage)
	r.GET("/test", TestPage)
}
