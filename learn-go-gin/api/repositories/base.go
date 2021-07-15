package repositories

import (
	"hotels-api/infrastructure"
)

// BaseRepository base repo struct
type BaseRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewBaseRepository(
	db infrastructure.Database,
	logger infrastructure.Logger,
) BaseRepository {
	return BaseRepository{
		db:     db,
		logger: logger,
	}
}

func (b BaseRepository) Create(m interface{}) error {
	b.logger.Info("create of base repository")
	return b.db.Create(m).Error
}

func (b BaseRepository) Update(m interface{}) error {
	return b.db.Save(m).Error
}

func (b BaseRepository) GetOne(m interface{}, pk interface{}) (err error) {
	return b.db.Where("id = ?", pk).First(m).Error
}

func (b BaseRepository) Delete(m interface{}, pk interface{}) error {
	return b.db.Delete(m, pk).Error
}
