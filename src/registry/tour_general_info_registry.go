package registry

import (
	"github.com/takatakayone/vrtravel_user_backend/src/interface_adapter/handler"
	ir "github.com/takatakayone/vrtravel_user_backend/src/interface_adapter/repository"
	ip "github.com/takatakayone/vrtravel_user_backend/src/interface_adapter/presenter"
	"github.com/takatakayone/vrtravel_user_backend/src/usecase/interactor"
	"github.com/takatakayone/vrtravel_user_backend/src/usecase/presenter"
	"github.com/takatakayone/vrtravel_user_backend/src/usecase/repository"
)

func (r *registry) NewTourGeneralInfoHandler() handler.TourGeneralInfoHandler {
	return handler.NewTourGeneralInfoHandler(r.NewTourGeneralInfoInteractor())
}

func (r *registry) NewTourGeneralInfoInteractor() interactor.TourGeneralInfoInteractor {
	return interactor.NewTourGeneralInfoInteractor(r.NewTourGeneralInfoRepository(),r.NewTourGeneralInfoPresenter())
}

func (r *registry) NewTourGeneralInfoRepository() repository.TourGeneralInfoRepository {
	return ir.NewTourGeneralInfoRepository()
}

func (r *registry) NewTourGeneralInfoPresenter() presenter.TourGeneralInfoPresenter {
	return ip.NewTourGeneralInfoPresenter()
}