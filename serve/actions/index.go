package actions

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func SystemInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"time": time.Now().Unix(),
	})
}
