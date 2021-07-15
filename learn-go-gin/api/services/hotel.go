package services

import (
	"hotels-api/api/repositories"
	"hotels-api/models"
)

type HotelService struct {
	repo repositories.HotelRepository
}

func NewHotelService(repo repositories.HotelRepository) HotelService {
	return HotelService{
		repo: repo,
	}
}

func (h HotelService) CreateHotel(hotel *models.Hotel) error {
	return h.repo.Create(hotel)
}
