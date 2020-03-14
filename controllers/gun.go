package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-1/models"
	log "github.com/sirupsen/logrus"
)

func RegisterGun(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg models.Gun

	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
	}

	data := models.Gun{
		Gun: msg.Gun,
	}

	message := models.RegisterGun(data)

	// log.Println(message)

	var dataresponse []StatusMessage

	dataresponse = append(dataresponse, StatusMessage{Message: message})

	if message == "Magazine number already registered" {

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

func LoadMagazine(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg models.LoadGun

	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
	}

	data := models.LoadGun{
		Magazine: msg.Magazine,
		Gun:      msg.Gun,
	}

	message := models.LoadMagazine(data)

	// log.Println(message)

	var dataresponse []StatusMessage

	dataresponse = append(dataresponse, StatusMessage{Message: message})

	if message == "Magazine already filled" {

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

func FireGun(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg models.Gun

	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
	}

	data := models.Gun{
		Gun: msg.Gun,
	}

	message := models.FireGun(data)

	// log.Println(message)

	var dataresponse []StatusMessage

	dataresponse = append(dataresponse, StatusMessage{Message: message})

	if message == "Gun Cannot Fired changing magazine try again" {

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
