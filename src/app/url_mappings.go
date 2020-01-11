package app

import (
	"github.com/takatakayone/vrtravel_user_backend/src/domain/ping/dao"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/ping/repositories"
	dao2 "github.com/takatakayone/vrtravel_user_backend/src/domain/tour/dao"
	repositories2 "github.com/takatakayone/vrtravel_user_backend/src/domain/tour/repositories"
	"github.com/takatakayone/vrtravel_user_backend/src/http/handler"
	"github.com/takatakayone/vrtravel_user_backend/src/usecases/ping"
	"github.com/takatakayone/vrtravel_user_backend/src/usecases/tour"
)

func urlMappings() {
	pingHandler := handler.NewPingHandler(ping.NewPingUsecase(repositories.NewPingRepository(dao.NewPingDao())))
	tourHandler := handler.NewTourHandler(tour.NewTourUseCase(repositories2.NewTourRepo(dao2.NewTourDao())))

	router.GET("/ping/:id", pingHandler.HandlePing)
	router.GET("/tour/:country", tourHandler.HandleGetTourGeneralInfoByCountry)
}
