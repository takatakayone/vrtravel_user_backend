package repositories

import (
	"github.com/takatakayone/vrtravel_user_backend/src/domain/ping/dao"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/ping/entities"
)


type PingRepositoryInterface interface {
	GetPingById(int64) (*entities.PingEntity, error)
}

type pingRepository struct {
	PingDao dao.PingDaoInterface
}

func NewPingRepository (dao dao.PingDaoInterface) PingRepositoryInterface{
	return &pingRepository{
		PingDao: dao,
	}
}

func (p *pingRepository) GetPingById(id int64) (*entities.PingEntity, error) {

	result, err := p.PingDao.GetPingById(id)
	if err != nil{
		return nil, err
	}

	return result, nil
}