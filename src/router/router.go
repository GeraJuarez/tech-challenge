package router

import (
	"tech-challenge/controller"

	"github.com/gorilla/mux"
)

func Start(c controller.AppController) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter()

	addHellokRouter(c, api)
	addCallAPI(c, api)

	return router
}

func addHellokRouter(c controller.AppController, api *mux.Router) {
	kvstore := api.PathPrefix("/hello").Subrouter()
	kvstore.HandleFunc("", c.Hello.GetValue).Methods("GET")
}

func addCallAPI(c controller.AppController, api *mux.Router) {
	extAPI := api.PathPrefix("/v1/pokemon").Subrouter()

	extAPI.HandleFunc("/{name}", c.ExtAPI.GetValue).Methods("GET")
}
