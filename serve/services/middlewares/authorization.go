package middlewares

import (
	"errors"
	"gotodo/serve/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

var noAuthPaths = []string{
	"/api/SignUp",
	"/api/SignIn",
}

func Authorization(c *gin.Context) {
	for _, path := range noAuthPaths {
		if path == c.Request.URL.Path {
			return
		}
	}

	token := c.GetHeader("Access-Token")
	if token == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("no token"))
	}

	_, err := core.Detoken(token)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
	}
}
