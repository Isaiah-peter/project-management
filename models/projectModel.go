package models

import "github.com/jinzhu/gorm"

type AddProjctModel struct {
	gorm.Model
	UserId      int64  `json:"user_id"`
	ProjectName string `json:"project_name"`
	Task        []Task
}

func (p *AddProjctModel) CreateProject() *AddProjctModel {
	db.NewRecord(p)
	db.Create(p)
	return p
}
