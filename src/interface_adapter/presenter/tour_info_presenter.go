package presenter

import "github.com/takatakayone/vrtravel_user_backend/src/domain/entity"

type tourInfoPresenter struct {

}

type TourInfoPresenter interface {
	ResponseTourInfos([]*entity.TourInfo,) ([]*entity.TourInfo, error)
}

func NewTourInfoPresenter() TourInfoPresenter{
	return &tourInfoPresenter{}
}

func (tp *tourInfoPresenter) ResponseTourInfos(ti []*entity.TourInfo,) ([]*entity.TourInfo, error) {
	return ti, nil
}


