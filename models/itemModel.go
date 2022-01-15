package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	TaskId   int64  `json:"task_id"`
	ItemName string `json:"item_name"`
}

func (i *Item) CreateTaskItem() *Item {
	db.Create(i)
	return i
}
