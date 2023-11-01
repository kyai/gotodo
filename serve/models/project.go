package models

type Project struct {
	Title  string   `json:"title"`
	Topics []*Topic `json:"topics" gorm:"-"`
	Model
}

type Topic struct {
	ProjectId int64   `json:"project_id"`
	Title     string  `json:"title"`
	Tasks     []*Task `json:"tasks" gorm:"-"`
	Model
}

type Task struct {
	ProjectId int64  `json:"project_id"`
	TopicId   int64  `json:"topic_id"`
	Caption   string `json:"caption"`
	Content   string `json:"content"`
	Model
}

func GetProjects() ([]*Project, error) {
	list := make([]*Project, 0)
	err := DB.Find(&list).Error
	return list, err
}

func LoadProject(id int64) (project *Project, err error) {
	project = &Project{}
	if err = DB.First(project, "id = ?", id).Error; err != nil {
		return
	}

	topics := make([]*Topic, 0)
	if err = DB.Where("project_id = ?", id).Find(&topics).Error; err != nil {
		return
	}

	tasks := make([]*Task, 0)
	if err = DB.Where("project_id = ?", id).Find(&tasks).Error; err != nil {
		return
	}

	// 组合数据
	topicMap := make(map[int64]*Topic)
	for _, topic := range topics {
		topicMap[topic.Id] = topic
	}
	for _, task := range tasks {
		topicMap[task.TopicId].Tasks = append(topicMap[task.TopicId].Tasks, task)
	}

	project.Topics = topics
	return
}
