package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {

}
func AddTag(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "created"})
}

func EditTag(c *gin.Context) {
}

func DeleteTag(c *gin.Context) {
}
