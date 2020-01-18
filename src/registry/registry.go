package registry

import "github.com/takatakayone/vrtravel_user_backend/src/interface_adapter/handler"

type registry struct {

}

type Registry interface {
	NewAppHandler() handler.AppHandler
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppHandler() handler.AppHandler {
	return r.NewTourGeneralInfoHandler()
}

