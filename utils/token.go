package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	fmt.Println(strArr)
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
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

func ValidToken(r *http.Request) (*jwt.Token, string) {
	token, err := VerifyToken(r)
	msg := ""
	if err != nil {
		msg := "verification fail"
		return nil, msg
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		msg := "token is not valid"
		return nil, msg
	}
	return token, msg
}

func UseToken(w http.ResponseWriter, r *http.Request) jwt.MapClaims {
	token, msg := ValidToken(r)
	if msg != "" {
		w.Write([]byte(msg))
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		panic(ok)
	}
	fmt.Print(claim)
	return claim
}
