package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project-management/models"
	"project-management/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	newProject models.AddProjctModel
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	newProject := &models.AddProjctModel{}
	Id, err := strconv.ParseInt(fmt.Sprintf("%f", token["UserId"]), 0, 0)
	if err != nil {
		log.Panic(err)
	}
	if err := json.NewDecoder(r.Body).Decode(newProject); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to decode json"))
	}
	newProject.UserId = Id
	u := newProject.CreateProject()
	res, errM := json.Marshal(u)
	if errM != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error while sending data back"))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	project := &models.AddProjctModel{}
	u := db.Preload("Task").Find(project).Value
	res, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("fail to send result"))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProjectById(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	newProject := &models.AddProjctModel{}
	muxid := mux.Vars(r)
	id, err := strconv.ParseInt(muxid["id"], 0, 0)
	if err != nil {
		log.Panic(err)
	}

	u := db.Where("ID =?", id).Find(&newProject).Value
	res, errRes := json.Marshal(u)
	if errRes != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("fail to return value"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProjectByUserId(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	userId, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserId"]), 0, 0)
	if err != nil {
		log.Panic(err)
	}
	u := db.Where("user_id=?", userId).Find(&newProject).Value
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
