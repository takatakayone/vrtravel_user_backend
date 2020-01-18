package entity

import "time"

type TourGeneralInfo struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	TourSchedules []TourSchedule `json:"tour_schedules"`
	TouristSpots []TouristSpot `json:"tourist_spots"`
}

type TourSchedule struct {
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
}

type TouristSpot struct {
	Name string `json:"name"`
	Location Location `json:"location"`
}

type Location struct {
	Country string `json:"country"`
	State string `json:"state"`
	City string `json:"city"`
}

