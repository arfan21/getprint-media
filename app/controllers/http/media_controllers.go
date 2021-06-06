package http

import (
	"fmt"
	"net/http"

	helpers2 "github.com/arfan21/getprint-media/app/helpers"
	services "github.com/arfan21/getprint-media/app/services"
	"github.com/labstack/echo/v4"
)

type MediaControllers interface {
	Create(c echo.Context) error
	Delete(c echo.Context) error
}

type mediaControllers struct {
	mediaSrv services.MediaServices
}

func NewMediaControllers(mediaSrv services.MediaServices) MediaControllers {
	return &mediaControllers{mediaSrv}
}

func (ctrl *mediaControllers) Create(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers2.Response("error", err.Error(), nil))
	}

	data, err := ctrl.mediaSrv.Create(c.Request().Context(), file)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers2.Response("error", err.Error(), nil))
	}
	fmt.Println(data)

	return c.JSON(http.StatusOK, helpers2.Response("success", nil, data))
}

func (ctrl *mediaControllers) Delete(c echo.Context) error {
	url := c.QueryParam("url")

	err := ctrl.mediaSrv.Delete(c.Request().Context(), url)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers2.Response("error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helpers2.Response("success", nil, nil))
}
