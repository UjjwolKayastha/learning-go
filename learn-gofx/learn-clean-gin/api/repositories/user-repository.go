package repositories

import (
	"github.com/ujjwolkayastha/learn-clean-gin/lib"
	"github.com/ujjwolkayastha/learn-clean-gin/models"
)

type UserRepository struct {
	db lib.Database
	logger lib.Logger
}

func NewUserRepository(db lib.Database, logger lib.Logger) UserRepository {
	return UserRepository{
		db: db,
		logger: logger,
	}
}

//get all users
func (r UserRepository) GetAll() (users []models.User, err error) {
	return users, r.db.DB.Find(&users).Error
}

// Save user
func (r UserRepository) Save(user models.User) (models.User, error) {
	return user, r.db.DB.Create(&user).Error
}

// Update updates user
func (r UserRepository) Update(user models.User) (models.User, error) {
	return user, r.db.DB.Save(&user).Error
}

// GetOne gets ont user
func (r UserRepository) GetOne(id uint) (user models.User, err error) {
	return user, r.db.DB.Where("id = ?", id).First(&user).Error
}

// Delete deletes the row of data
func (r UserRepository) Delete(id uint) error {
	return r.db.DB.Where("id = ?", id).Delete(&models.User{}).Error
}