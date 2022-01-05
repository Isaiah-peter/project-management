package models

import (
	"log"
	"project-management/database"
	"project-management/utils"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique" form:"email" binding:"required"`
	Password  string ` json:"password" binding:"required,min=6"`
}

var db = database.GetDB()

func (u *User) CreateUser() *User {
	hashpassword, err := utils.GeneratePassword(u.Password)
	if err != nil {
		log.Panic(err)
	}
	u.Password = hashpassword
	db.NewRecord(u)
	db.Create(u)
	return u
}
