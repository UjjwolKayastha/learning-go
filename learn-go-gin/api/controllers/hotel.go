package controllers

import (
	"fmt"
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

// HandleGetAllHotels get all hotels
func (h HotelController) HandleGetAllHotels() gin.HandlerFunc {
	return func(c *gin.Context) {
		hotels, count, err := h.service.GetAllHotels()
		if err != nil {
			h.logger.Error("error getting hotels ", err)
			responses.ErrorJSON(c, http.StatusInternalServerError, err)
			return
		}

		responses.JSONCount(c, http.StatusOK, hotels, int(count))
	}
}

// HandleGetOneHotel get one hotel
func (h HotelController) HandleGetOneHotel() gin.HandlerFunc {
	return func(c *gin.Context) {

		hotelID := c.Param("hotelID")
		hotel := models.Hotel{}

		if err := h.service.GetOneHotel(&hotel, hotelID); err != nil {
			h.logger.Error("error getting hotel ", err)
			responses.ErrorJSON(c, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(c, http.StatusOK, hotel)
	}
}

func (h HotelController) HandleUpdateHotel() gin.HandlerFunc {
	return func(c *gin.Context) {
		hotelID := c.Param("hotelID")
		hotel := models.Hotel{}

		fmt.Println("FROM CONT", hotelID, hotel)

		if err := h.service.GetOneHotel(&hotel, hotelID); err != nil {
			h.logger.Error("error getting hotel ", err)
			responses.ErrorJSON(c, http.StatusInternalServerError, err)
			return
		}

		if err := c.ShouldBindJSON(&hotel); err != nil {
			responses.ErrorJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := h.service.UpdateHotel(hotelID, &hotel); err != nil {
			h.logger.Error("error updaing hotel", err)
			responses.ErrorJSON(c, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(c, http.StatusOK, "hotel updated")
	}
}
