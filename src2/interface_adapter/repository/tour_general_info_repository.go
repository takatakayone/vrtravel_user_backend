package repository

import (
	"database/sql"
	"errors"
	"github.com/takatakayone/vrtravel_user_backend/src/datasources/mysql"
	"github.com/takatakayone/vrtravel_user_backend/src2/domain/datamodel"
	"github.com/takatakayone/vrtravel_user_backend/src2/domain/entity"
	"log"
)

const(
	queryTourGeneralInfo =
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

type tourGeneralInfoRepository struct {
	//db *sql.DB
}

type TourGenralInfoRepository interface {
	FindTourGeneralInfosByCountry(string) ([]*entity.TourGeneralInfo, error)
}

func NewTourGeneralInfoRepository(db *sql.DB) TourGenralInfoRepository {
	return &tourGeneralInfoRepository{}
}


// datasourceがmysqlでもElasticsearchでもこのInterfaceを実装すればOKで入れ替え簡単。
func (tr *tourGeneralInfoRepository) FindTourGeneralInfosByCountry(country string) ([]*entity.TourGeneralInfo, error) {
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

	queryResults := make([]datamodel.TourGeneralInfo, 0)
	for rows.Next() {
		var tourGenralInfo datamodel.TourGeneralInfo
		if err := rows.Scan(&tourGenralInfo.Id, &tourGenralInfo.Title, &tourGenralInfo.StartTime, &tourGenralInfo.EndTime ,&tourGenralInfo.Country, &tourGenralInfo.State, &tourGenralInfo.City, &tourGenralInfo.Name); err != nil{
			log.Println("error when tyring to scan tourgeneral info", err)
			return nil, err
		}

		queryResults = append(queryResults, tourGenralInfo)
	}

	if len(queryResults) == 0 {
		err := errors.New("not_found")
		log.Println("error when tyring to tour general info", err)
		return nil ,err
	}

	result :=  make([] *entity.TourGeneralInfo, 0)
	tourIds := make([]int64, 0)
	for _, tour := range queryResults {
		var tourGeneralEntity entity.TourGeneralInfo

		// もっとなんとかしたい...なんか冗長。。
		//すでにtourIdがある場合はすでにあるtourGeneralEntityのTouristSpotsにappendする形。ない場合にはtourEntityを作る感じ。
		if contains(tourIds, tour.Id) {
			tourGeneralInfoEntity := findTourGeneralEntityById(result, tour.Id)
			tourGeneralInfoEntity.TourSchedules = append(tourGeneralInfoEntity.TourSchedules, tour.MappingTourSchedule())
			tourGeneralInfoEntity.TouristSpots = append(tourGeneralInfoEntity.TouristSpots, tour.MappingTouristSpot())
		} else {
			tourIds = append(tourIds, tour.Id)
			tourGeneralEntity = tour.Mapping()
			result = append(result, &tourGeneralEntity)
		}
	}

	return result, nil
}

// 下のふたつの関数化はもっと抽象化してutilsとして使えるようにした方が良さそう。
func findTourGeneralEntityById(elements [] *entity.TourGeneralInfo, tourId int64) (*entity.TourGeneralInfo) {
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