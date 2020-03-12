package controllers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type StatusMessage struct {
	Message string `json:"message"`
}
type Response struct {
	Code   int         `json:"code"`   // 200 , 400
	Status string      `json:"status"` // "Ok" "Error"
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var dataresponse []StatusMessage

	dataresponse = append(dataresponse, StatusMessage{Message: "Welcome to Home controller"})

	response := &Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   dataresponse,
	}

	urlsJson, _ := json.Marshal(response)
	log.Println(urlsJson)
	//log.Println(Bloguser)

	w.Write(urlsJson)

}
