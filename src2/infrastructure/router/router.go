package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takatakayone/vrtravel_user_backend/src2/infrastructure/datastore/mysql"
	"github.com/takatakayone/vrtravel_user_backend/src2/interface_adapter/handler"
	"github.com/takatakayone/vrtravel_user_backend/src2/interface_adapter/presenter"
	"github.com/takatakayone/vrtravel_user_backend/src2/interface_adapter/repository"
	"github.com/takatakayone/vrtravel_user_backend/src2/usecase/interactor"
)

var(
	router = gin.Default()
)

func StartApplication() {
	urlMappings()

	router.Run(":8080")
}

func urlMappings()  {
	tourGeneralInfoHandler := handler.NewTourGeneralInfoHandler(interactor.NewTourGeneralInfoInteractor(repository.NewTourGeneralInfoRepository(mysql.UsersDb), presenter.NewTourGeneralInfoPresenter()))
	router.GET("/tour/:country", tourGeneralInfoHandler.GetTourGeneralInfoByCountry)
}