package datamodel

import (
	"github.com/takatakayone/vrtravel_user_backend/src/domain/entity"
	"time"
)

type TourInfo struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Country string `json:"country"`
	State string `json:"state"`
	City string `json:"city"`
	Name string `json:"name"`
}

func (tg *TourInfo) Mapping() entity.TourInfo {
	return entity.TourInfo{
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

func (tg *TourInfo) MappingTouristSpot () entity.TouristSpot {
	return entity.TouristSpot{
		Name:     tg.Name,
		Location: entity.Location{
			Country: tg.Country,
			State:   tg.State,
			City:    tg.City,
		},
	}
}

func (tg *TourInfo) MappingTourSchedule () entity.TourSchedule {
	return entity.TourSchedule{
		StartTime: tg.StartTime,
		EndTime:   tg.EndTime,
	}
}
