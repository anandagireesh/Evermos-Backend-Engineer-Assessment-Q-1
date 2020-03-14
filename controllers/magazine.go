package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-1/models"
	log "github.com/sirupsen/logrus"
)

func RegisterMagazine(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg models.Magazine

	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
	}

	data := models.Magazine{
		Magazine: msg.Magazine,
	}

	message := models.RegisterMagazine(data)

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

// fill magazine with bullets

func FillMagazine(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg models.FillMagazine

	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
	}

	data := models.FillMagazine{
		Magazine: msg.Magazine,
		Bullets:  msg.Bullets,
	}

	message := models.FillMAgazine(data)

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

func VerifyMag(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg models.Magazine

	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
	}

	data := models.Magazine{
		Magazine: msg.Magazine,
	}

	message := models.VerifyMagazine(data)

	log.Println(message)

	var dataresponse []StatusMessage

	dataresponse = append(dataresponse, StatusMessage{Message: message})

	if message == "Magazine not verified" {

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
