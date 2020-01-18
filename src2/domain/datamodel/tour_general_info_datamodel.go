package datamodel

import (
	"github.com/takatakayone/vrtravel_user_backend/src2/domain/entity"
	"time"
)

type TourGeneralInfo struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Country string `json:"country"`
	State string `json:"state"`
	City string `json:"city"`
	Name string `json:"name"`
}

func (tg *TourGeneralInfo) Mapping() entity.TourGeneralInfo {
	return entity.TourGeneralInfo{
		Id:				 tg.Id,
		Title:           tg.Title,
		TourSchedules:   []entity.TourSchedule {
			tg.MappingTourSchedule(),
		},
		TouristSpots:    []entity.TouristSpot{
			tg.MappingTouristSpot(),
		},
	}
}

func (tg *TourGeneralInfo) MappingTouristSpot () entity.TouristSpot {
	return entity.TouristSpot{
		Name:     tg.Name,
		Location: entity.Location{
			Country: tg.Country,
			State:   tg.State,
			City:    tg.City,
		},
	}
}

func (tg *TourGeneralInfo) MappingTourSchedule () entity.TourSchedule {
	return entity.TourSchedule{
		StartTime: tg.StartTime,
		EndTime:   tg.EndTime,
	}
}
