package mysql

import (
	"context"

	models2 "github.com/arfan21/getprint-media/app/models"
	"github.com/arfan21/getprint-media/configs/mysql"
)

type MediaRepository interface {
	Create(ctx context.Context, data *models2.Media) error
	Delete(ctx context.Context, url string) error
	GetByURL(ctx context.Context, url string) (*models2.Media, error)
}

type mediaRepository struct {
	db mysql.Client
}

func NewMediaRepository(db mysql.Client) MediaRepository {
	return &mediaRepository{db}
}

func (repo *mediaRepository) Create(ctx context.Context, data *models2.Media) error {
	return repo.db.Conn().WithContext(ctx).Create(data).Error
}

func (repo *mediaRepository) GetByURL(ctx context.Context, url string) (*models2.Media, error) {
	media := new(models2.Media)
	err := repo.db.Conn().WithContext(ctx).Where("url = ?", url).First(media).Error
	if err != nil {
		return nil, err
	}

	return media, nil
}

func (repo *mediaRepository) Delete(ctx context.Context, url string) error {
	return repo.db.Conn().WithContext(ctx).Unscoped().Where("url = ?", url).Delete(&models2.Media{}).Error
}
