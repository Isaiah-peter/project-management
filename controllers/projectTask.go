package controllers

import (
	"encoding/json"
	"net/http"
	"project-management/models"
	"project-management/utils"
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
