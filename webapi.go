package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/yanndr/webapi/config"
	"github.com/yanndr/webapi/router"
)

func main() {
	config.Init()

	router := router.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(config.Get().Port, n)
}
