package models

type Model struct {
	Id      int64 `gorm:"primaryKey" json:"id"`
	Created int64 `gorm:"autoCreateTime" json:"created"`
	Updated int64 `gorm:"autoUpdateTime" json:"updated"`
	Deleted int64 `gorm:"autoDeleteTime" json:"-"`
}
