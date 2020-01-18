package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/takatakayone/vrtravel_user_backend/src2/usecase/interactor"
	"net/http"
)

type tourGeneralInfoHandler struct {
	tourGeneralInfoInteractor interactor.TourGeneralInfoInteractor
}

type TourGeneralInfoHandler interface {
	GetTourGeneralInfoByCountry(*gin.Context)
}

func NewTourGeneralInfoHandler(interactor interactor.TourGeneralInfoInteractor) TourGeneralInfoHandler{
	return &tourGeneralInfoHandler{tourGeneralInfoInteractor:interactor}
}

func (th *tourGeneralInfoHandler) GetTourGeneralInfoByCountry(c *gin.Context) {
	country := c.Param("country")
	if country == "" {
		c.JSON(http.StatusBadRequest, "errrorr params")
		return
	}
	result, err := th.tourGeneralInfoInteractor.GetTourGeneralInfosByCountry(country)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}