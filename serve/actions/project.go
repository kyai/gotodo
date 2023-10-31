package actions

import (
	"gotodo/serve/models"

	"github.com/gin-gonic/gin"
)

func Projects(c *gin.Context) {
	list, err := models.GetProjects()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, list)
}

func AddProject(c *gin.Context) {
	var payload struct {
		Title string `json:"title"`
	}
	c.BindJSON(&payload)

	pro := &models.Project{
		Title: payload.Title,
	}

	if err := models.DB.Create(pro).Error; err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, pro.Id)
}

func PutProject(c *gin.Context) {}

func DelProject(c *gin.Context) {}
