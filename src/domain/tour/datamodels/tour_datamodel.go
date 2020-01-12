package datamodels

import (
	"github.com/takatakayone/vrtravel_user_backend/src/domain/tour/entities"
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

func (tg *TourGeneralInfo) Mapping() entities.TourGeneralInfo {
	return entities.TourGeneralInfo{
		Id:				 tg.Id,
		Title:           tg.Title,
		TourSchedules:   []entities.TourSchedule {
				tg.MappingTourSchedule(),
		},
		TouristSpots:    []entities.TouristSpot{
			tg.MappingTouristSpot(),
		},
	}
}

func (tg *TourGeneralInfo) MappingTouristSpot () entities.TouristSpot {
	return entities.TouristSpot{
		Name:     tg.Name,
		Location: entities.Location{
			Country: tg.Country,
			State:   tg.State,
			City:    tg.City,
		},
	}
}

func (tg *TourGeneralInfo) MappingTourSchedule () entities.TourSchedule {
	return entities.TourSchedule{
		StartTime: tg.StartTime,
		EndTime:   tg.EndTime,
	}
}