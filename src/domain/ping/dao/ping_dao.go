package dao

import (
	"github.com/takatakayone/vrtravel_user_backend/src/datasources/mysql"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/ping/entities"
	"log"
)

const(
	queryGetPingById = "SELECT * FROM pings WHERE id=?;"
)

type PingDaoInterface interface {
	GetPingById(id int64) (*entities.PingEntity, error)
}

type pingDao struct {

}

func NewPingDao() *pingDao{
	return &pingDao{}
}

func (p *pingDao) GetPingById(id int64) (entity *entities.PingEntity, err error){
	result := entities.PingEntity{}
	stmt, err := mysql.UsersDb.Prepare(queryGetPingById)

	if err != nil{
		log.Println("stmt error when trying to get ping by id")
		return nil, err
	}
	defer stmt.Close()
	queryResult := stmt.QueryRow(id)
	if err := queryResult.Scan(&result.Id, &result.Title, &result.Body, &result.CreatedAt); err != nil{
		return nil, err
	}

	return &result,nil
}