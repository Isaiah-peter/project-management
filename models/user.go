package models

import (
	"log"
	"project-management/database"
	"project-management/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique" form:"email" binding:"required"`
	Password  string ` json:"password" binding:"required,min=6"`
}

type Token struct {
	UserId int
	Email  string
	jwt.StandardClaims
}

var db *gorm.DB

func Init() {
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate(&User{}, &AddProjctModel{}, &Task{}, &Item{})
}

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

func GenerateToken(password, email string) map[string]interface{} {
	user := &User{}

	if err := db.Where("email =?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}

	expireAt := time.Now().Add(time.Minute * 100000).Unix()

	msg, errf := utils.Checkpassword(user.Password, password)
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		var resp = map[string]interface{}{"status": false, "message": msg}
		return resp
	}

	tk := &Token{
		UserId: int(user.ID),
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		panic(err)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user
	return resp

}
