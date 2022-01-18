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
	projectTaskItem models.Item
	decodeMsg       string = "error while decoding json"
)

func CreateProjectTaskItem(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(w, r)
	if err := json.NewDecoder(r.Body).Decode(&projectTask); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(decodeMsg))
	}
	u := projectTaskItem.CreateTaskItem()
	res, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("fail to return"))
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(res)
}

func GetProjectTaskItem(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(w, r)
	db.Find(&projectTaskItem)
	res, err := json.Marshal(&projectTaskItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("fail to send result"))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProjectTaskItemById(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(w, r)
	var projectItem = models.Item{}
	var Id = mux.Vars(r)
	var projectItemId, err = strconv.ParseInt(Id["id"], 0, 0)
	if err != nil {
		log.Panic(err)
	}

	db.Find(&projectItem, projectItemId)
	res, _ := json.Marshal(&projectItem)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
