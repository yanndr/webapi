package router

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/yanndr/webapi/controller"
	"github.com/yanndr/webapi/middleware"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	SetEntityRoutes(router)
	SetEntityAuthRoutes(router)
	return router
}

func SetEntityRoutes(router *mux.Router) {
	router.Handle("/entity",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controller.GetEntities),
		)).Methods("GET")

}

func SetEntityAuthRoutes(router *mux.Router) {
	router.Handle("/authentity",
		negroni.New(
			negroni.HandlerFunc(middleware.RequireTokenAuthentication),
			negroni.HandlerFunc(controller.GetEntities),
		)).Methods("GET")

}

func SetAuthenticationRoutes(router *mux.Router) {
	router.HandleFunc("/token-auth", controller.Login).Methods("POST")
}
