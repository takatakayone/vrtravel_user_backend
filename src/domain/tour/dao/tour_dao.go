package dao

import (
	"github.com/takatakayone/vrtravel_user_backend/src/datasources/mysql"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/tour/datamodels"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/tour/entities"
	"log"
)

const(
	queryTourGeneralInfo =
	`SELECT  tour.id, tour.title, lo.country, lo.state, lo.city,  ts.name
	FROM locations AS lo 
	JOIN tourist_spots AS ts
	ON lo.id = ts.location_id
	JOIN tour_schedules AS tsh
	ON ts.id = tsh.tourist_spot_id
	JOIN tours AS tour
	ON tour.id = tsh.tour_id
	WHERE lo.country=?;`
)

type TourDaoInterface interface {
	FetchTourInfosByCountry(string) (*[] *entities.TourInfo, error)
}

type tourDao struct {

}

func NewTourDao() TourDaoInterface{
	return &tourDao{}
}

func (d *tourDao) FetchTourInfosByCountry(country string) (*[] *entities.TourInfo, error) {

	stmt, err := mysql.UsersDb.Prepare(queryTourGeneralInfo)

	if err != nil{
		log.Println("stmt error when trying to get GeneralTourInfoByCountry")
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(country)
	if err != nil {
		log.Println("error when trying to query generalTourInfo", err)
		return nil, err
	}
	defer rows.Close()

	queryResults := make([]datamodels.TourGeneralInfo, 0)
	for rows.Next() {
		var tourGenralInfo datamodels.TourGeneralInfo
		if err := rows.Scan(&tourGenralInfo.Id, &tourGenralInfo.Title, &tourGenralInfo.Country, &tourGenralInfo.State, &tourGenralInfo.City, &tourGenralInfo.Name); err != nil{
			log.Println("error when tyring to scan tourgeneral info", err)
			return nil, err
		}

		queryResults = append(queryResults, tourGenralInfo)
	}

	if len(queryResults) == 0 {
		log.Println("error when tyring to tour general info", err)
		return nil ,err
	}

	result :=  make([] *entities.TourInfo, 0)
	tourIds := make([]int64, 0)
	for _, tour := range queryResults {
		var tourGeneralEntity entities.TourInfo

		// もっとなんとかしたい...なんか冗長。。
		//すでにtourIdがある場合はすでにあるtourGeneralEntityのTouristSpotsにappendする形。ない場合にはtourEntityを作る感じ。
		if contains(tourIds, tour.Id) {
			tourGeneralInfoEntity := findElementById(result, tour.Id)
			tourGeneralInfoEntity.TouristSpots = append(tourGeneralInfoEntity.TouristSpots, tour.MappingTouristSpot())
		} else {
			tourIds = append(tourIds, tour.Id)
			tourGeneralEntity = tour.Mapping()
			result = append(result, &tourGeneralEntity)
		}
	}

	return &result, nil
}

func findElementById(elements [] *entities.TourInfo, tourId int64) (*entities.TourInfo) {
	for _, element := range elements {
		if tourId == element.Id {
			return element
		}
	}
	return nil
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}