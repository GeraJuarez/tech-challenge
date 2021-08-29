package registry

import (
	"reflect"
	"tech-challenge/controller"
	"tech-challenge/usecase/interactor"
	"testing"
)

func TestCreateExtAPIInteractor(t *testing.T) {
	r := registry{}
	inter := r.NewExtAPIInteractor()

	expected := reflect.TypeOf(interactor.NewExtApiInteractor())
	result := reflect.TypeOf(inter)

	if expected != result {
		t.Errorf("wrong type: got %v want %v", result, expected)
	}
}

func TestCreateExtAPIController(t *testing.T) {
	r := registry{}
	cont := r.NewExtAPIController()

	expected := reflect.TypeOf(controller.NewExtApiController(r.NewExtAPIInteractor()))
	result := reflect.TypeOf(cont)

	if expected != result {
		t.Errorf("wrong type: got %v want %v", result, expected)
	}
}
