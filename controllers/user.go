package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project-management/database"
	"project-management/models"
	"project-management/utils"
	"strconv"
)

var db = database.GetDB()

func GetUser(w http.ResponseWriter, r *http.Request) {
	var newUser []models.User
	u := db.Find(&newUser).Limit(10)
	res, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error while trying to get all user"))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)

}
func NewUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	err := json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	u := newUser.CreateUser()
	res, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("fail to convert json to golang understandable language"))
	}
	u := models.GenerateToken(newUser.Password, newUser.Email)
	res, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("login fail"))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	token := utils.UseToken(r)
	ID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		log.Fatalln(err)
	}

	u := db.Where("ID =?", ID).Find(user)
	res, errr := json.Marshal(u)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error while changing to json"))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}
