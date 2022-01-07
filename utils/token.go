package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

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
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok && token.Valid {
		log.Panic(ok)
		msg = fmt.Sprintln("token is not correct or bad token")
	}
	if claim.VerifyExpiresAt(time.Now().Local().Unix(), true) == false {
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
