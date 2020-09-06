package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Home(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "raw")
}

func NowPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
	})
}

func TestPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
