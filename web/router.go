package web

import (
	"github.com/gin-gonic/gin"
	"github.com/twocucao/thanos/web/apps/admin"
	"github.com/twocucao/thanos/web/apps/misc"
	"github.com/twocucao/thanos/web/apps/pages"
	"github.com/twocucao/thanos/web/apps/user"
)

func InitRouter(r *gin.Engine) {
	user.InitRoutes(r)
	admin.InitRoutes(r)
	misc.InitRoutes(r)
	pages.InitRoutes(r)
}
