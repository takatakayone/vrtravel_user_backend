package entities

type TourInfo struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	TouristSpots []TouristSpot `json:"tourist_spots"`
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
