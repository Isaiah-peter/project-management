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
	newProject models.Project
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(w, r)
	newProject := &models.Project{}
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
	utils.UseToken(w, r)
	project := &models.Project{}
	db.Preload("Task").Find(project)
	res, err := json.Marshal(&project)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("fail to send result"))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProjectById(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(w, r)
	newProject := &models.Project{}
	muxid := mux.Vars(r)
	id, err := strconv.ParseInt(muxid["id"], 0, 0)
	if err != nil {
		log.Panic(err)
	}

	db.Where("ID =?", id).Preload("Task").Find(&newProject)
	res, errRes := json.Marshal(&newProject)
	if errRes != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("fail to return value"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProjectByUserId(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(w, r)
	userId, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserId"]), 0, 0)
	if err != nil {
		log.Panic(err)
	}
	db.Where("user_id=?", userId).Find(&newProject)
	res, _ := json.Marshal(&newProject)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(w, r)
	project := &models.Project{}
	var Id = mux.Vars(r)
	projectId, err := strconv.ParseInt(Id["id"], 0, 0)
	if err != nil {
		log.Panic(err)
	}

	if err := json.NewDecoder(r.Body).Decode(project); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic(err)
	}

	u := db.Where("ID=?", projectId).Find(project)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	u.Save(&project)
	w.Write([]byte("successfully update"))
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	project := &models.Project{}
	utils.UseToken(w, r)
	var id = mux.Vars(r)
	projectId, err := strconv.ParseInt(id["id"], 0, 0)
	if err != nil {
		log.Panic(err)
	}

	db.Where("ID=?", projectId).Delete(&project)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfully delete project"))
}
