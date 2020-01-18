package presenter

import "github.com/takatakayone/vrtravel_user_backend/src2/domain/entity"

type tourGeneralInfoPresenter struct {

}

type TourGeneralInfoPresenter interface {
	ResponseTourGeneralInfos([]*entity.TourGeneralInfo,) ([]*entity.TourGeneralInfo, error)
}

func NewTourGeneralInfoPresenter() TourGeneralInfoPresenter{
	return &tourGeneralInfoPresenter{}
}

func (tp *tourGeneralInfoPresenter) ResponseTourGeneralInfos(ti []*entity.TourGeneralInfo,) ([]*entity.TourGeneralInfo, error) {
	return ti, nil
}


