package models

type Project struct {
	Id    int64 `gorm:"primaryKey"`
	Title string
}

func GetProjects() ([]*Project, error) {
	list := make([]*Project, 0)
	err := DB.Find(&list).Error
	return list, err
}
