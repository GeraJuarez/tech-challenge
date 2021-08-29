package interactor_test

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"tech-challenge/model"
	"tech-challenge/usecase/interactor"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestPokemonSuccess(t *testing.T) {
	inter := interactor.NewExtApiInteractor()
	expected := model.Pokemon{Order: 1, Name: "bulbasaur", Weight: 69, Height: 7}
	result, _ := inter.Get("1")

	if expected != result {
		t.Errorf("wrong type: got %v want %v", result, expected)
	}
}

func TestPokemonFail(t *testing.T) {
	inter := interactor.NewExtApiInteractor()
	_, err := inter.Get("0")

	if !errors.Is(err, interactor.ErrorPkmnNotFound) {
		t.Errorf("wrong type: got %v want %v", err, interactor.ErrorPkmnNotFound)
	}
}
