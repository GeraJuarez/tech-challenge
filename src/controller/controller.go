package controller

import (
	"encoding/json"
	"errors"
	"example/cloud-app/store/usecase/interactor"
	"net/http"

	"github.com/gorilla/mux"
)

type AppController struct {
	Hello  HelloController
	ExtAPI ExtApiController
}

// Hello
type HelloController interface {
	GetValue(w http.ResponseWriter, r *http.Request)
}

type helloController struct{}

func NewHelloController() HelloController {
	return &helloController{}
}

func (c *helloController) GetValue(w http.ResponseWriter, r *http.Request) {
	value := "Hello, World."

	w.Write([]byte(value))
}

// External API
type ExtApiController interface {
	GetValue(w http.ResponseWriter, r *http.Request)
}

type extApiController struct {
	extApiInteractor interactor.ExtApiInteractor
}

func NewExtApiController(interactor interactor.ExtApiInteractor) ExtApiController {
	return &extApiController{interactor}
}

func (c *extApiController) GetValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	value, err := c.extApiInteractor.Get(name)

	if errors.Is(err, interactor.ErrorPkmnNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if errors.Is(err, interactor.ErrorPkmnServerError) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
