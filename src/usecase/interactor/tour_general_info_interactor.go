package interactor

import (
	"fmt"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/entity"
	"github.com/takatakayone/vrtravel_user_backend/src/usecase/presenter"
	"github.com/takatakayone/vrtravel_user_backend/src/usecase/repository"
)

type tourGeneralInfoInteractor struct {
	TourGeneralInfoRepo repository.TourGeneralInfoRepository
	TourGeneralInfoPresenter presenter.TourGeneralInfoPresenter
}

type TourGeneralInfoInteractor interface {
	GetTourGeneralInfosByCountry(country string) ([]*entity.TourGeneralInfo, error)
}

func NewTourGeneralInfoInteractor(r repository.TourGeneralInfoRepository, p presenter.TourGeneralInfoPresenter) TourGeneralInfoInteractor {
	return &tourGeneralInfoInteractor{r, p}
}

func (ti *tourGeneralInfoInteractor) GetTourGeneralInfosByCountry(country string) ([]*entity.TourGeneralInfo, error) {
	tourGeneralInfos, err := ti.TourGeneralInfoRepo.FindTourGeneralInfosByCountry(country)

	fmt.Println(tourGeneralInfos, err)

	if err != nil {
		return nil, err
	}

	return ti.TourGeneralInfoPresenter.ResponseTourGeneralInfos(tourGeneralInfos)
}