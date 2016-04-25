package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yanndr/webapi/config"
	"github.com/yanndr/webapi/handler"
)

func main() {
	serverConfig, err := config.LoadServerConfiguration()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/entity", handler.GetEntities).Methods("GET")
	log.Fatal(http.ListenAndServe(serverConfig.Port, router))
}
