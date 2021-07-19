package repositories

import (
	"hotels-api/models"
)

// HotelRepository hotel repository
type HotelRepository struct {
	BaseRepository
}

// NewHotelRepository creates new booking repository
func NewHotelRepository(
	base BaseRepository,
) HotelRepository {
	return HotelRepository{
		BaseRepository: base,
	}
}

// GetAll get all the hotels addeed
func (b HotelRepository) GetAll() (hotels []models.Hotel, count int64, err error) {
	return hotels, count, b.db.Model(&models.Hotel{}).Find(&hotels).Count(&count).Error
}

func (b HotelRepository) UpdateOne(id string, hotel *models.Hotel) error {
	return b.db.Model(&models.Hotel{}).Where("id =? ", id).Updates(map[string]interface{}{
		"name":        hotel.Name,
		"location":    hotel.Location,
		"description": hotel.Description,
	}).Error
}

// Delete deletes the row of data
func (b HotelRepository) Delete(id string) error {
	return b.db.Where("id = ? ", id).Delete(&models.Hotel{}).Error
}
