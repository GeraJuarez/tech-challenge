package interactor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tech-challenge/model"
)

var ErrorPkmnNotFound = errors.New("pokemon not found")
var ErrorPkmnServerError = errors.New("unable to call pkmn API")

type extApiInteractor struct{}

type ExtApiInteractor interface {
	Get(name string) (model.Pokemon, error)
}

func NewExtApiInteractor() ExtApiInteractor {
	return &extApiInteractor{}
}

func (api *extApiInteractor) Get(name string) (model.Pokemon, error) {
	var pokemon model.Pokemon
	requestURL := fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%s", name)

	response, err := http.Get(requestURL)
	if err != nil {
		log.Fatal(err)
		return pokemon, ErrorPkmnServerError
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return pokemon, ErrorPkmnServerError
	}

	log.Printf("Pokemon API Status response %d", response.StatusCode)

	if response.StatusCode == http.StatusNotFound {
		return pokemon, ErrorPkmnNotFound
	} else {
		var pokemon model.Pokemon
		json.Unmarshal(responseData, &pokemon)

		return pokemon, nil
	}
}
