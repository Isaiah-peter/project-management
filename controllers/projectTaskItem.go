package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project-management/models"
	"project-management/utils"
)

var (
	projectTaskItem models.Item
	decodeMsg       string = "error while decoding json"
)

func CreateProjectTaskItem(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	if err := json.NewDecoder(r.Body).Decode(projectTask); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(decodeMsg))
	}
	u := projectTaskItem.CreateTaskItem()
	res, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("fail to return"))
		w.Write([]byte(fmt.Sprintln("%v", err)))
	}
	w.Write(res)
}
