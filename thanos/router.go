package thanos

import (
	"github.com/gin-gonic/gin"
	"github.com/twocucao/thanos/thanos/apps/admin"
	"github.com/twocucao/thanos/thanos/apps/misc"
	"github.com/twocucao/thanos/thanos/apps/pages"
	"github.com/twocucao/thanos/thanos/apps/user"
)

func InitRouter(r *gin.Engine) {
	user.InitRoutes(r)
	admin.InitRoutes(r)
	misc.InitRoutes(r)
	pages.InitRoutes(r)
}
