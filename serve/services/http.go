package services

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
		api.POST("SystemInfo", actions.SystemInfo)
		api.POST("SystemConfigs", actions.SystemConfigs)
		api.POST("GetSystemConfig", actions.GetSystemConfig)
		api.POST("SetSystemConfig", actions.SetSystemConfig)
	}

	r.Run(flags.Host + ":" + flags.Port)
}
