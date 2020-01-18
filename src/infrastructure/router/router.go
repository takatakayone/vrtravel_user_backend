package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takatakayone/vrtravel_user_backend/src/interface_adapter/handler"
)

var(
	router = gin.Default()
)

func NewRouter(handler handler.AppHandler) {
	urlMappings(handler)
	router.Run(":8080")
}