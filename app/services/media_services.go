package services

import (
	"context"
	"mime/multipart"

	"github.com/arfan21/getprint-media/app/models"
	"github.com/arfan21/getprint-media/app/repository/dropbox"
	repo "github.com/arfan21/getprint-media/app/repository/mysql"
)

type MediaServices interface {
	Create(ctx context.Context, file *multipart.FileHeader) (*models.Media, error)
	Delete(ctx context.Context, url string) error
}

type mediaServices struct {
	mediaRepo repo.MediaRepository
	dbx       dropbox.Dropbox
}

func NewMediaServices(mediaRepo repo.MediaRepository) MediaServices {
	dbx := dropbox.NewDropboxRepository()

	return &mediaServices{mediaRepo, dbx}
}

func (srv *mediaServices) Create(ctx context.Context, file *multipart.FileHeader) (*models.Media, error) {
	path, err := srv.dbx.Uploader(file)
	if err != nil {
		return nil, err
	}

	url, err := srv.dbx.CreateSharedLink(path)
	if err != nil {
		srv.dbx.Delete(path)
		return nil, err
	}

	data := new(models.Media)
	data.Path = path
	data.Url = url

	err = srv.mediaRepo.Create(ctx, data)
	if err != nil {
		srv.dbx.Delete(path)
		return nil, err
	}

	return data, nil
}

func (srv *mediaServices) Delete(ctx context.Context, url string) error {
	data, err := srv.mediaRepo.GetByURL(ctx, url)
	if err != nil {
		return err
	}

	err = srv.dbx.Delete(data.Path)
	if err != nil {
		return err
	}

	err = srv.mediaRepo.Delete(ctx, url)
	if err != nil {
		return err
	}

	return nil
}
