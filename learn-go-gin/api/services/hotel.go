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

func (h HotelService) GetOneHotel(hotel *models.Hotel, id string) error {
	return h.repo.GetOne(hotel, id)
}

func (h HotelService) GetAllHotels() ([]models.Hotel, int64, error) {
	return h.repo.GetAll()
}

func (h HotelService) UpdateHotel(id string, hotel *models.Hotel) error {
	return h.repo.UpdateOne(id, hotel)
}

func (h HotelService) DeleteOneHotel(hotel *models.Hotel, id string) error {
	return h.repo.Delete(id)
}
