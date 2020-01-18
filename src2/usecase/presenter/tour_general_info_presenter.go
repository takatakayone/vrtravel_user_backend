package presenter

import "github.com/takatakayone/vrtravel_user_backend/src2/domain/entity"

type TourGeneralInfoPresenter interface {
	ResponseTourGeneralInfos([]*entity.TourGeneralInfo,) ([]*entity.TourGeneralInfo, error)
}