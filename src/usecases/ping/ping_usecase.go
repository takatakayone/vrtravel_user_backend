package ping

import (
	"github.com/takatakayone/vrtravel_user_backend/src/domain/ping/entities"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/ping/repositories"
)

type PingUsecaseInterface interface {
	GetPongById(int64) (*entities.PingEntity, error)
}

type pingUsecase struct {
	PingRepo repositories.PingRepositoryInterface
}

func NewPingUsecase(repo repositories.PingRepositoryInterface) PingUsecaseInterface {
	return &pingUsecase{
		PingRepo: repo,
	}
}

func (p *pingUsecase) GetPongById(id int64) (entity *entities.PingEntity, err error) {

	result, err := p.PingRepo.GetPingById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
