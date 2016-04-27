package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/yanndr/webapi/model"
)

func GetEntities(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	res.Header().Set("Content-Type", "application/json")

	var entities = [...]*model.Entity{
		&model.Entity{Nkey: "123"},
		&model.Entity{Nkey: "123"},
	}

	outgoingJSON, error := json.Marshal(entities)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}
