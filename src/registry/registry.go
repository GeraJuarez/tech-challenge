package registry

import (
	"tech-challenge/controller"
	"tech-challenge/usecase/interactor"
)

type registry struct{}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Hello:  r.NewHelloController(),
		ExtAPI: r.NewExtAPIController(),
	}
}

func (r *registry) NewHelloController() controller.HelloController {
	return controller.NewHelloController()
}

func (r *registry) NewExtAPIController() controller.ExtApiController {
	return controller.NewExtApiController(r.NewExtAPIInteractor())
}

func (r *registry) NewExtAPIInteractor() interactor.ExtApiInteractor {
	return interactor.NewExtApiInteractor()
}
