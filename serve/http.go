package serve

import (
	"gotodo/serve/actions"
	"gotodo/serve/flags"

	"github.com/gin-gonic/gin"
)

func HttpHandler() {
	r := gin.Default()
	r.LoadHTMLGlob(flags.View + "/*")

	r.GET("/", actions.Index)

	api := r.Group("/api")
	{
		api.POST("/sysinfo", actions.SystemInfo)
	}

	r.Run(flags.Host + ":" + flags.Port)
}
