package repository

import "github.com/takatakayone/vrtravel_user_backend/src/domain/entity"

type TourGeneralInfoRepository interface {
	FindTourGeneralInfosByCountry(string) ([]*entity.TourGeneralInfo, error)
}