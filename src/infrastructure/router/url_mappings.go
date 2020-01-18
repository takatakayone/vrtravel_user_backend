package router

import (
	"github.com/takatakayone/vrtravel_user_backend/src/interface_adapter/handler"
)

func urlMappings(handler handler.AppHandler)  {
	router.GET("/tour/:country", handler.GetTourGeneralInfoByCountry)
}
