package controller

import (
	"encoding/json"
	"net/http"

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
