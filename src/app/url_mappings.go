package app

import (
	"github.com/takatakayone/vrtravel_user_backend/src/http/handler"
)

func urlMappings() {
	pingHandler := handler.NewPingHandler()

	router.GET("/ping/:id", pingHandler.HandlePing)
}
