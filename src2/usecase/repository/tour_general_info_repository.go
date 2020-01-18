package repository

import "github.com/takatakayone/vrtravel_user_backend/src2/domain/entity"

type TourGeneralInfoRepository interface {
	FindTourGeneralInfosByCountry(string) ([]*entity.TourGeneralInfo, error)
}