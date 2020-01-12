package tour

import (
	"github.com/takatakayone/vrtravel_user_backend/src/domain/tour/entities"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/tour/repositories"
	"log"
)

type TourUsecaseInterface interface {
	GetTourGeneralInfosByCountry(country string) (*[] *entities.TourGeneralInfo, error)
}

type tourUsecase struct {
	repo repositories.TourRepostitoryInterface
}

func NewTourUseCase(repo repositories.TourRepostitoryInterface) TourUsecaseInterface {
	return &tourUsecase{repo:repo}
}

func (u *tourUsecase) GetTourGeneralInfosByCountry(country string)(*[] *entities.TourGeneralInfo, error) {
	result, err := u.repo.GetTourGeneralInfosByCountry(country)
	if err != nil{
		log.Println("error usecase")
		return nil, err
	}
	return result, err
}
