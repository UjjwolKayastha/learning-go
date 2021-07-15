package controllers

import (
	"hotels-api/api/responses"
	"hotels-api/api/services"
	"hotels-api/infrastructure"
	"hotels-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HotelController Hotel controller
type HotelController struct {
	logger  infrastructure.Logger
	service services.HotelService
}

// NewHotelController creates new owner controller
func NewHotelController(
	logger infrastructure.Logger,
	service services.HotelService,
) HotelController {
	return HotelController{
		service: service,
		logger:  logger,
	}
}

// HandleCreateHotel creates hotel
func (h HotelController) HandleCreateHotel() gin.HandlerFunc {
	return func(c *gin.Context) {

		hotel := models.Hotel{}
		if err := c.ShouldBindJSON(&hotel); err != nil {
			h.logger.Error("error binding input: ", err)
			responses.ErrorJSON(c, http.StatusInternalServerError, err)
			return
		}

		if err := h.service.CreateHotel(&hotel); err != nil {
			h.logger.Error("error creating hotel: ", err)
			responses.ErrorJSON(c, http.StatusInternalServerError, err)
			return
		}

		responses.SuccessJSON(c, http.StatusCreated, "hotel created")

	}
}
