package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	ProjectId int64  `json:"project_id"`
	TaskName  string `json:"task_name"`
	Items     []Item `json:"items"`
}