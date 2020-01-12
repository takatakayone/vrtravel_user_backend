package tour

import (
	"github.com/takatakayone/vrtravel_user_backend/src/domain/tour/entities"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/tour/repositories"
	"log"
)

type TourUsecaseInterface interface {
	GetTourInfosByCountry(country string) (*[] *entities.TourInfo, error)
}

type tourUsecase struct {
	repo repositories.TourRepostitoryInterface
}

func NewTourUseCase(repo repositories.TourRepostitoryInterface) TourUsecaseInterface {
	return &tourUsecase{repo:repo}
}

func (u *tourUsecase) GetTourInfosByCountry(country string)(*[] *entities.TourInfo, error) {
	result, err := u.repo.GetTourInfosByCountry(country)
	if err != nil{
		log.Println("error usecase")
		return nil, err
	}
	return result, err
}
