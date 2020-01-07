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

func NewPingRepository (daoInterface dao.PingDaoInterface) *pingRepository{
	return &pingRepository{}
}

func (p *pingRepository) GetPingById(id int64) (*entities.PingEntity, error) {
	p.PingDao.GetPingById(id)

	return nil, nil
}