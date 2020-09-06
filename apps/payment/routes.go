package payment

import (
	"github.com/gin-gonic/gin"
	"zelda/apps/payment/api"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/", api.Home)
	r.GET("/raw", api.NowPage)
	r.GET("/test", api.TestPage)
}
