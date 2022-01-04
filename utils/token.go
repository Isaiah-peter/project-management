package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"project-management/database"
	"project-management/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var db = database.GetDB()

type Token struct {
	UserId int
	Email  string
	jwt.StandardClaims
}

func GenerateToken(password, email string) map[string]interface{} {
	user := &models.User{}

	if err := db.Where("email =?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}

	expireAt := time.Now().Add(time.Minute * 100000).Unix()

	msg, errf := Checkpassword(user.Password, password)
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		var resp = map[string]interface{}{"status": false, "message": msg}
		return resp
	}

	tk := &Token{
		UserId: int(user.Model.ID),
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

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("token")
	return token
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method: %v", t.Header["alg"])
		}
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ValidateToken(r *http.Request) (string, error) {
	token, err := VerifyToken(r)
	msg := ""
	if err != nil {
		log.Panic(err)
		msg = fmt.Sprintln("token verification fail")
	}
	claim, ok := token.Claims.(*Token)
	if !ok && token.Valid {
		log.Panic(ok)
		msg = fmt.Sprintln("token is not correct or bad token")
	}
	if claim.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token exprired")
	}
	return msg, nil
}

func UseToken(r *http.Request) jwt.MapClaims {
	token, err := VerifyToken(r)
	if err != nil {
		panic(err)
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		panic(ok)
	}
	return claim
}
