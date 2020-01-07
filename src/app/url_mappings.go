package app

import (
	"github.com/takatakayone/vrtravel_user_backend/src/domain/ping/dao"
	"github.com/takatakayone/vrtravel_user_backend/src/domain/ping/repositories"
	"github.com/takatakayone/vrtravel_user_backend/src/http/handler"
	"github.com/takatakayone/vrtravel_user_backend/src/usecases/ping"
)

func urlMappings() {
	pingHandler := handler.NewPingHandler(ping.NewPingUsecase(repositories.NewPingRepository(dao.NewPingDao())))

	router.GET("/ping/:id", pingHandler.HandlePing)
}
