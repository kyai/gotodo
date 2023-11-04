package models

import (
	"gorm.io/plugin/soft_delete"
)

type Model struct {
	Id      int64                 `gorm:"primaryKey" json:"id"`
	Created int64                 `gorm:"autoCreateTime" json:"created"`
	Updated int64                 `gorm:"autoUpdateTime" json:"updated"`
	Deleted soft_delete.DeletedAt `json:"-"`
}
