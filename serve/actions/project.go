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

func GetProject(c *gin.Context) {
	var payload struct {
		Id int64 `json:"id"`
	}
	c.BindJSON(&payload)

	pro, err := models.LoadProject(payload.Id)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, pro)
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

func AddTopic(c *gin.Context) {
	var payload struct {
		ProjectId int64  `json:"project_id"`
		Title     string `json:"title"`
	}
	c.BindJSON(&payload)

	if err := models.DB.Create(&models.Topic{
		ProjectId: payload.ProjectId,
		Title:     payload.Title,
	}).Error; err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, nil)
}

func AddTask(c *gin.Context) {
	var payload struct {
		ProjectId int64  `json:"project_id"`
		TopicId   int64  `json:"topic_id"`
		Caption   string `json:"caption"`
		Content   string `json:"content"`
	}
	c.BindJSON(&payload)

	if err := models.DB.Create(&models.Task{
		ProjectId: payload.ProjectId,
		TopicId:   payload.TopicId,
		Caption:   payload.Caption,
		Content:   payload.Content,
	}).Error; err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, nil)
}
