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

// GetOne gets one hotel
func (b HotelRepository) GetOne(m *models.Hotel, id string) (err error) {
	return b.db.Where("id = ? ", id).Preload("Owner").First(m).Error
}

// Delete deletes the row of data
func (b HotelRepository) Delete(id, ownerID string) error {
	return b.db.Where("id = ? AND owner_id = ?", id, ownerID).Delete(&models.Hotel{}).Error
}

// GetAll get all the hotels addeed
func (b HotelRepository) GetAll() (hotels []models.Hotel, count int64, err error) {
	return hotels, count, b.db.Model(&models.Hotel{}).Find(&hotels).Count(&count).Error
}
