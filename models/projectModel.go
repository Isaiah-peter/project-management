package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	UserId      int64  `json:"user_id"`
	ProjectName string `json:"project_name"`
	Task        []Task
}

func (p *Project) CreateProject() *Project {
	db.Create(p)
	return p
}
