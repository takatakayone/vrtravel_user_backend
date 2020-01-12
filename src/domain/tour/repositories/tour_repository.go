package repositories

import (
	"github.com/takatakayone/vrtravel_user_backend/src/domain/tour/dao"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/tour/entities"
	"log"
)

type TourRepostitoryInterface interface {
	GetTourInfosByCountry(country string) (*[]*entities.TourInfo, error)
}

type tourRepo struct {
	dao dao.TourDaoInterface
}

func NewTourRepo(dao dao.TourDaoInterface) TourRepostitoryInterface {
	return &tourRepo{dao: dao}
}

func (r *tourRepo) GetTourInfosByCountry(country string)(*[] *entities.TourInfo, error) {
	result, err := r.dao.FetchTourInfosByCountry(country)
	if err != nil {
		log.Println("error when trying to fetch general info dao")
		return nil, err
	}
	return result, nil
}