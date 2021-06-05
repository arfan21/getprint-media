package media

import (
	"context"

	"github.com/arfan21/getprint-media/models"
	"gorm.io/gorm"
)

type MediaRepository interface {
	Create(ctx context.Context, data *models.Media) error
	Delete(ctx context.Context, url string) error
	GetByURL(ctx context.Context, url string) (*models.Media, error)
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

func (repo *mediaRepository) GetByURL(ctx context.Context, url string) (*models.Media, error) {
	media := new(models.Media)
	err := repo.db.WithContext(ctx).Where("url = ?", url).First(media).Error
	if err != nil {
		return nil, err
	}

	return media, nil
}

func (repo *mediaRepository) Delete(ctx context.Context, url string) error {
	return repo.db.WithContext(ctx).Unscoped().Where("url = ?", url).Delete(&models.Media{}).Error
}
