package presenter

import "github.com/takatakayone/vrtravel_user_backend/src/domain/entity"

type TourInfoPresenter interface {
	ResponseTourInfos([]*entity.TourInfo,) ([]*entity.TourInfo, error)
}