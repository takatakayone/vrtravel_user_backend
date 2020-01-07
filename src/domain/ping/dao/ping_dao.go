package dao

import (
	"github.com/takatakayone/vrtravel_user_backend/src/datasources/mysql"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/ping/datamodels"
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
	pingData := datamodels.Ping{}
	stmt, err := mysql.UsersDb.Prepare(queryGetPingById)

	if err != nil{
		log.Println("stmt error when trying to get ping by id")
		return nil, err
	}
	defer stmt.Close()
	queryResult := stmt.QueryRow(id)
	if err := queryResult.Scan(&pingData.Id, &pingData.Title, &pingData.Body, &pingData.CreatedAt); err != nil{
		return nil, err
	}

	result := pingData.Mapping()

	return &result,nil
}