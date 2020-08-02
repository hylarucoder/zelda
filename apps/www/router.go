package www

import (
	"github.com/gin-gonic/gin"
	"github.com/twocucao/zelda/apps/admin"
	"github.com/twocucao/zelda/apps/misc"
	"github.com/twocucao/zelda/apps/pages"
	"github.com/twocucao/zelda/apps/user"
)

func InitRouter(r *gin.Engine) {
	user.InitRoutes(r)
	admin.InitRoutes(r)
	misc.InitRoutes(r)
	pages.InitRoutes(r)
}
