package www

import (
	"github.com/gin-gonic/gin"
	"zelda/apps/admin"
	"zelda/apps/misc"
	"zelda/apps/pages"
	"zelda/apps/user"
)

func InitRouter(r *gin.Engine) {
	user.InitRoutes(r)
	admin.InitRoutes(r)
	misc.InitRoutes(r)
	pages.InitRoutes(r)
}
