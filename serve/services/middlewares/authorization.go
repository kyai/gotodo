package middlewares

import (
	"gotodo/serve/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

var noAuthPaths = []string{
	"/api/Login",
}

func Authorization(c *gin.Context) {
	for _, path := range noAuthPaths {
		if path == c.Request.URL.Path {
			return
		}
	}

	token := c.GetHeader("Access-Token")
	if token == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	_, err := core.Detoken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
