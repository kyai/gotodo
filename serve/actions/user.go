package actions

import (
	"errors"
	"gotodo/serve/core"
	"gotodo/serve/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 注册
func SignUp(c *gin.Context) {
	// TODO 是否开启注册

	var payload struct {
		Username string `json:"username"`
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}
	c.BindJSON(&payload)

	if err := models.DB.Create(&models.User{
		Username: payload.Username,
		Nickname: payload.Nickname,
		Password: core.MD5(payload.Password + core.AppKey),
	}).Error; err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "OK")
}

// 登录
func SignIn(c *gin.Context) {
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.BindJSON(&payload)

	user := &models.User{}
	if err := models.DB.Where("username = ?", payload.Username).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.String(500, "用户名或密码错误")
		} else {
			c.String(500, err.Error())
		}
		return
	}

	if user.Password != core.MD5(payload.Password+core.AppKey) {
		c.String(500, "用户名或密码错误")
		return
	}

	token, err := core.Entoken(user.Id)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, map[string]interface{}{
		"token": token,
		"user":  user,
	})
}

func Profile(c *gin.Context) {
	user, err := models.GetUser(c.GetInt64("UID"))
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, user)
}
