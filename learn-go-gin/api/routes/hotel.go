package routes

import (
	"hotels-api/api/controllers"
	"hotels-api/infrastructure"
)

// HotelRoutes hotel routes
type HotelRoutes struct {
	logger     infrastructure.Logger
	router     infrastructure.Router
	controller controllers.HotelController
}

// NewHotelRoutes creates new hotel routes
func NewHotelRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	controller controllers.HotelController,
) HotelRoutes {
	return HotelRoutes{
		router:     router,
		controller: controller,
		logger:     logger,
	}
}

// Setup sets up routes
func (h HotelRoutes) Setup() {
	h.logger.Info("setting up Hotel routes")
	hotel := h.router.Group("/hotel")
	{
		hotel.POST("", h.controller.HandleCreateHotel())
	}

}
