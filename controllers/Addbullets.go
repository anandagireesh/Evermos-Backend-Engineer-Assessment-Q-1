package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-1/models"
	log "github.com/sirupsen/logrus"
)

// register bullets to database controller

func AddBullet(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg models.Bullets

	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
	}

	data := models.Bullets{
		Bullets: msg.Bullets,
	}

	message := models.AddBullet(data)

	// log.Println(message)

	var dataresponse []StatusMessage

	dataresponse = append(dataresponse, StatusMessage{Message: message})

	if message == "Bullet numbered already registered" {

		response := &Response{
			Code:   http.StatusNotAcceptable,
			Status: "Not Acceptable",
			Data:   dataresponse,
		}
		urlsJson, _ := json.Marshal(response)
		log.Println(urlsJson)

		w.Write(urlsJson)

	} else {

		response := &Response{
			Code:   http.StatusOK,
			Status: "ok",
			Data:   dataresponse,
		}
		urlsJson, _ := json.Marshal(response)
		log.Println(urlsJson)

		w.Write(urlsJson)

	}
}
