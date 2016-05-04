package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yanndr/webapi/mapper"
	"github.com/yanndr/webapi/model"
	"github.com/yanndr/webapi/service"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(model.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := service.Login(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func Register(w http.ResponseWriter, r *http.Request) {
	registerRequest := model.User{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&registerRequest)
	_, err := mapper.CreateUser(registerRequest.Username, registerRequest.Password)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
