package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"project-management/models"
	"project-management/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	projectTask models.Task
)

func CreateProjectTask(w http.ResponseWriter, g http.Response, r *http.Request) {
	utils.UseToken(r)
	if err := json.NewDecoder(r.Body).Decode(&projectTask); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error while decoding"))
	}

	u := projectTask.CreateTask()
	res, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to convert to json"))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}

func GetProjectTask(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	db.Preload("Items").Find(&projectTask)
	res, err := json.Marshal(&projectTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("fail to send result"))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProjectTaskByProjectId(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	projectId := mux.Vars(r)
	id, err := strconv.ParseInt(projectId["id"], 0, 0)
	if err != nil {
		log.Panic(err)
	}

	db.Where("project_id = ?", id).Preload("Items").First(&projectTask)
	res, err1 := json.Marshal(&projectTask)
	if err1 != nil {
		log.Panic(err1)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProjectTask(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)

}
