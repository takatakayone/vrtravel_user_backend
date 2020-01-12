package datamodels

import (
	"github.com/takatakayone/vrtravel_user_backend/src/domain/tour/entities"
)

type TourGeneralInfo struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Country string `json:"country"`
	State string `json:"state"`
	City string `json:"city"`
	Name string `json:"name"`
}

func (tg *TourGeneralInfo) Mapping() entities.TourInfo {

	return entities.TourInfo{
		Id:				 tg.Id,
		Title:           tg.Title,
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





//type TourSchedule struct {
//	Id int64 `json:"id"`
//	Content string `json:"content"`
//	TourId int64 `json:"tour_id"`
//	StartTime time.Time `json:"start_time"`
//	EndTime time.Time `json:"end_time"`
//	TouristSpotId int64 `json:"tourist_spot_id"`
//	CreatedAt time.Time `json:"created_at"`
//	UpdatedAt time.Time `json:"updated_at"`
//}
//
//type TouristSpot struct {
//	Id int64 `json:"id"`
//	Name string `json:"name"`
//	CreatedAt time.Time `json:"created_at"`
//	UpdatedAt time.Time `json:"updated_at"`
//}
//
//type Location struct {
//	Id int64 `json:"id"`
//	Country string `json:"country"`
//	State string `json:"state"`
//	City string `json:"city"`
//	CreatedAt time.Time `json:"created_at"`
//	UpdatedAt time.Time `json:"updated_at"`
//}