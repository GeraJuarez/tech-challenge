package controller_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"tech-challenge/model"
	"tech-challenge/registry"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestHello(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	registry := registry.NewRegistry()
	rr := httptest.NewRecorder()
	handler := mux.NewRouter()
	handler.HandleFunc("/hello", registry.NewAppController().Hello.GetValue)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Hello, World.`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPokemonSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/pokemon/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	registry := registry.NewRegistry()
	rr := httptest.NewRecorder()
	handler := mux.NewRouter()
	handler.HandleFunc("/v1/pokemon/{name}", registry.NewAppController().ExtAPI.GetValue)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := model.Pokemon{Order: 1, Name: "bulbasaur", Weight: 69, Height: 7}
	var result model.Pokemon
	responseData, _ := ioutil.ReadAll(rr.Body)
	json.Unmarshal(responseData, &result)
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestPokemonByNameSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/pokemon/bulbasaur", nil)
	if err != nil {
		t.Fatal(err)
	}

	registry := registry.NewRegistry()
	rr := httptest.NewRecorder()
	handler := mux.NewRouter()
	handler.HandleFunc("/v1/pokemon/{name}", registry.NewAppController().ExtAPI.GetValue)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := model.Pokemon{Order: 1, Name: "bulbasaur", Weight: 69, Height: 7}
	var result model.Pokemon
	responseData, _ := ioutil.ReadAll(rr.Body)
	json.Unmarshal(responseData, &result)
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPokemonFailure(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/pokemon/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	registry := registry.NewRegistry()
	rr := httptest.NewRecorder()
	handler := mux.NewRouter()
	handler.HandleFunc("/v1/pokemon/{name}", registry.NewAppController().ExtAPI.GetValue)
	handler.ServeHTTP(rr, req)
	expected := http.StatusNotFound
	if status := rr.Code; status != expected {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expected)
	}
}
