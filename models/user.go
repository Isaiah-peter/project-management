package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
