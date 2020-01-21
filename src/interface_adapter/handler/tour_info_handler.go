package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/takatakayone/vrtravel_user_backend/src/usecase/interactor"
	"net/http"
)

type tourInfoHandler struct {
	tourInfoInteractor interactor.TourInfoInteractor
}

type TourInfoHandler interface {
	GetTourInfoByCountry(*gin.Context)
}

func NewTourGeneralInfoHandler(interactor interactor.TourInfoInteractor) TourInfoHandler{
	return &tourInfoHandler{tourInfoInteractor:interactor}
}

func (th *tourInfoHandler) GetTourInfoByCountry(c *gin.Context) {
	country := c.Param("country")
	if country == "" {
		c.JSON(http.StatusBadRequest, "errrorr params")
		return
	}
	result, err := th.tourInfoInteractor.GetTourInfosByCountry(country)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}