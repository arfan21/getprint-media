package media

import (
	"fmt"
	"net/http"

	"github.com/arfan21/getprint-media/helpers"
	_mediaRepo "github.com/arfan21/getprint-media/repository/mysql/media"
	_mediaSrv "github.com/arfan21/getprint-media/services/media"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MediaControllers interface {
	Routes(route *echo.Echo)
}

type mediaControllers struct {
	mediaSrv _mediaSrv.MediaServices
}

func NewMediaControllers(db *gorm.DB) MediaControllers {
	mediaRepo := _mediaRepo.NewMediaRepository(db)
	mediaSrv := _mediaSrv.NewMediaServices(mediaRepo)

	return &mediaControllers{mediaSrv}
}

func (ctrl *mediaControllers) Create(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response("error", err.Error(), nil))
	}

	data, err := ctrl.mediaSrv.Create(c.Request().Context(), file)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response("error", err.Error(), nil))
	}
	fmt.Println(data)

	return c.JSON(http.StatusOK, helpers.Response("success", nil, data))
}

func (ctrl *mediaControllers) Delete(c echo.Context) error {
	url := c.QueryParam("url")

	err := ctrl.mediaSrv.Delete(c.Request().Context(), url)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helpers.Response("success", nil, nil))
}
