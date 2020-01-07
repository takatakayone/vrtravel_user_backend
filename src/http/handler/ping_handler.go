package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/takatakayone/vrtravel_user_backend/src/usecases/ping"
	"net/http"
	"strconv"
)

type HandlePingInterface interface {
	HandlePing(*gin.Context)
}

type pingHandler struct {
	PingUsecase ping.PingUsecaseInterface
}

func NewPingHandler() *pingHandler  {
	return &pingHandler{}
}

func (p *pingHandler) HandlePing(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result, err := p.PingUsecase.GetPongById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "errorrr")
		return
	}
	c.JSON(http.StatusOK, result)
}
