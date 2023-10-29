package actions

import (
	"gotodo/serve/config"
	"time"

	"github.com/gin-gonic/gin"
)

func SystemInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"time": time.Now().Unix(),
	})
}

func SystemConfigs(c *gin.Context) {
	c.JSON(200, config.AllSettings())
}

func GetSystemConfig(c *gin.Context) {
	key := config.GetKey(c.Query("name"))
	if key == nil {
		c.String(500, "配置项无效")
		return
	}

	c.String(200, key.Get())
}

func SetSystemConfig(c *gin.Context) {
	var payload struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	c.BindJSON(&payload)

	key := config.GetKey(payload.Name)
	if key == nil {
		c.String(500, "配置项无效")
		return
	}
	if err := key.Set(payload.Value); err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "OK")
}
