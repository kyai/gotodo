package models

type model struct {
	Id int64 `gorm:"primaryKey;column:rowid"`
}
