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
	db.Find(&newUser).Limit(10)
	res, err := json.Marshal(&newUser)
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
	token := utils.UseToken(w, r)
	fmt.Println(token["UserID"])
	user := models.User{}
	ID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserId"]), 0, 0)
	fmt.Print(ID)
	if err != nil {
		fmt.Println("error while parsing")
		log.Fatalln(err)
	}

	db.Find(&user, ID)
	res, errr := json.Marshal(&user)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error while changing to json"))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(w, r)
	var user = models.User{}
	var fetchUser = models.User{}
	if err := json.NewDecoder(r.Body).Decode(*&user); err != nil {
		log.Panic(err)
	}

	if user.Email != "" {
		fetchUser.Email = user.Email
	}
	if user.FirstName != "" {
		fetchUser.FirstName = user.FirstName
	}

	if user.LastName != "" {
		fetchUser.LastName = user.LastName
	}
	if user.Password != "" {
		hash, err := utils.GeneratePassword(user.Password)
		if err != nil {
			log.Panic(err)
		}
		fetchUser.Password = hash
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("slusses"))

}
