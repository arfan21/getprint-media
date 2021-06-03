package media

import (
	"context"

	"github.com/arfan21/getprint-media/models"
	"gorm.io/gorm"
)

type MediaRepository interface {
	Create(ctx context.Context, data *models.Media) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*models.Media, error)
}

type mediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepository{db}
}

func (repo *mediaRepository) Create(ctx context.Context, data *models.Media) error {
	return repo.db.WithContext(ctx).Create(data).Error
}

func (repo *mediaRepository) GetByID(ctx context.Context, id uint) (*models.Media, error) {
	media := new(models.Media)
	err := repo.db.WithContext(ctx).First(media, id).Error
	if err != nil {
		return nil, err
	}

	return media, nil
}

func (repo *mediaRepository) Delete(ctx context.Context, id uint) error {
	return repo.db.WithContext(ctx).Unscoped().Delete(&models.Media{}, id).Error
}
