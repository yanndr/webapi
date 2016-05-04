package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	_ "github.com/lib/pq"
	"github.com/yanndr/webapi/config"
	"github.com/yanndr/webapi/database"
	"github.com/yanndr/webapi/router"
)

const (
	dbUser     = "postgres"
	dbPassword = "Antibes06"
	dbName     = "Ft"
)

func main() {
	config.Init()

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	var err error
	database.DBCon, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	defer database.DBCon.Close()

	router := router.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(config.Get().Port, n)
}
