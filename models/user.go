package models

import (
	"log"
	"project-management/database"
	"project-management/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func init() {
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&Item{})
}

func (u *User) CreateUser() *User {
	hashpassword, err := utils.GeneratePassword(u.Password)
	if err != nil {
		log.Panic(err)
	}
	u.Password = hashpassword
	db.Create(u)
	return u
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
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

	var resp = map[string]interface{}{"status": true, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user
	return resp

}
