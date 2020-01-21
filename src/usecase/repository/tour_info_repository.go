package repository

import "github.com/takatakayone/vrtravel_user_backend/src/domain/entity"

type TourInfoRepository interface {
	FetchTourInfosByCountry(string) ([]*entity.TourInfo, error)
}