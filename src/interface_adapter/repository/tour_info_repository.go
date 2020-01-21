package repository

import (
	"errors"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/datamodel"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/entity"
	"github.com/takatakayone/vrtravel_user_backend/src/infrastructure/datastore/mysql"
	"log"
)

const(
	queryTourInfo =
		`SELECT  tour.id, tour.title, tsh.start_time, tsh.end_time, lo.country, lo.state, lo.city,  ts.name
	FROM locations AS lo 
	JOIN tourist_spots AS ts
	ON lo.id = ts.location_id
	JOIN tour_schedules AS tsh
	ON ts.id = tsh.tourist_spot_id
	JOIN tours AS tour
	ON tour.id = tsh.tour_id
	WHERE lo.country=?;`
)

type tourInfoRepository struct {

}

type TourInfoRepository interface {
	FetchTourInfosByCountry(string) ([]*entity.TourInfo, error)
}

func NewTourInfoRepository() TourInfoRepository {
	return &tourInfoRepository{}
}


// datasourceがmysqlでもElasticsearchでもこのInterfaceを実装すればOKで入れ替え簡単。
func (tr *tourInfoRepository) FetchTourInfosByCountry(country string) ([]*entity.TourInfo, error) {
	stmt, err := mysql.UsersDb.Prepare(queryTourInfo)

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

	queryResults := make([]datamodel.TourInfo, 0)
	for rows.Next() {
		var tourInfo datamodel.TourInfo
		if err := rows.Scan(&tourInfo.Id, &tourInfo.Title, &tourInfo.StartTime, &tourInfo.EndTime ,&tourInfo.Country, &tourInfo.State, &tourInfo.City, &tourInfo.Name); err != nil{
			log.Println("error when tyring to scan tourgeneral info", err)
			return nil, err
		}

		queryResults = append(queryResults, tourInfo)
	}

	if len(queryResults) == 0 {
		err := errors.New("not_found")
		log.Println("error when tyring to tour general info", err)
		return nil ,err
	}

	result :=  make([] *entity.TourInfo, 0)
	tourIds := make([]int64, 0)
	for _, tour := range queryResults {
		var tourEntity entity.TourInfo

		// もっとなんとかしたい...なんか冗長。。
		//すでにtourIdがある場合はすでにあるtourGeneralEntityのTouristSpotsにappendする形。ない場合にはtourEntityを作る感じ。
		if contains(tourIds, tour.Id) {
			tourGeneralInfoEntity := findTourGeneralEntityById(result, tour.Id)
			tourGeneralInfoEntity.TourSchedules = append(tourGeneralInfoEntity.TourSchedules, tour.MappingTourSchedule())
			tourGeneralInfoEntity.TouristSpots = append(tourGeneralInfoEntity.TouristSpots, tour.MappingTouristSpot())
		} else {
			tourIds = append(tourIds, tour.Id)
			tourEntity = tour.Mapping()
			result = append(result, &tourEntity)
		}
	}

	return result, nil
}

// 下のふたつの関数化はもっと抽象化してutilsとして使えるようにした方が良さそう。
func findTourGeneralEntityById(elements [] *entity.TourInfo, tourId int64) (*entity.TourInfo) {
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