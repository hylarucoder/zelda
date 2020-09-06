package open

import (
	"github.com/gin-gonic/gin"
	"zelda/apps/official"
	"zelda/apps/user"
	"zelda/apps/payment"
	"zelda/apps/admin"
	"zelda/apps/operation"
)

func InitRouter(r *gin.Engine) {
	operation.InitRoutes(r)
	user.InitRoutes(r)
	payment.InitRoutes(r)
	admin.InitRoutes(r)
	official.InitRoutes(r)
}
