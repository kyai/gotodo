package services

import (
	"gotodo/serve/actions"
	"gotodo/serve/flags"
	"gotodo/serve/services/middlewares"

	"github.com/gin-gonic/gin"
)

func HttpHandler() {
	r := gin.Default()

	r.LoadHTMLGlob(flags.View + "/index.html")
	r.Static("/assets", flags.View+"/assets")
	r.StaticFile("/favicon.ico", flags.View+"/favicon.ico")

	r.GET("/", actions.Index)

	api := r.Group("/api")
	api.Use(middlewares.Authorization)
	{
		api.POST("SignUp", actions.SignUp)
		api.POST("SignIn", actions.SignIn)

		api.POST("Projects", actions.Projects)
		api.POST("AddProject", actions.AddProject)
		api.POST("PutProject", actions.PutProject)
		api.POST("DelProject", actions.DelProject)

		api.POST("SystemInfo", actions.SystemInfo)
		api.POST("SystemConfigs", actions.SystemConfigs)
		api.POST("GetSystemConfig", actions.GetSystemConfig)
		api.POST("SetSystemConfig", actions.SetSystemConfig)
	}

	r.Run(flags.Host + ":" + flags.Port)
}
