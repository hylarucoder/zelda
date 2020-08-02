package helpers

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/twocucao/zelda/conf"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * conf.PageSize
	}

	return result
}
