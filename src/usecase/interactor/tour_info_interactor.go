package interactor

import (
	"github.com/takatakayone/vrtravel_user_backend/src/domain/entity"
	"github.com/takatakayone/vrtravel_user_backend/src/usecase/presenter"
	"github.com/takatakayone/vrtravel_user_backend/src/usecase/repository"
)

type tourInfoInteractor struct {
	TourInfoRepo repository.TourInfoRepository
	TourInfoPresenter presenter.TourInfoPresenter
}

type TourInfoInteractor interface {
	GetTourInfosByCountry(country string) ([]*entity.TourInfo, error)
}

func NewTourInfoInteractor(r repository.TourInfoRepository, p presenter.TourInfoPresenter) TourInfoInteractor {
	return &tourInfoInteractor{r, p}
}

func (ti *tourInfoInteractor) GetTourInfosByCountry(country string) ([]*entity.TourInfo, error) {
	tourInfos, err := ti.TourInfoRepo.FetchTourInfosByCountry(country)

	if err != nil {
		return nil, err
	}

	return ti.TourInfoPresenter.ResponseTourInfos(tourInfos)
}