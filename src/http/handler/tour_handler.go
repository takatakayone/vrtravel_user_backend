package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/takatakayone/vrtravel_user_backend/src/usecases/tour"
	"net/http"
)

type TourHandlerInterface interface {
	HandleGetTourGeneralInfoByCountry(*gin.Context)
}

type tourHandler struct {
	TourUsecase tour.TourUsecaseInterface
}

func NewTourHandler(usecase tour.TourUsecaseInterface) TourHandlerInterface  {
	return &tourHandler{TourUsecase:usecase}
}

func (p *tourHandler) HandleGetTourGeneralInfoByCountry(c *gin.Context) {
	country := c.Param("country")
	if country == "" {
		c.JSON(http.StatusBadRequest, "errrorr params")
		return
	}
	result, err := p.TourUsecase.GetTourGeneralInfosByCountry(country)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}
