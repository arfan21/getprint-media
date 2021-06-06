package server

import (
	ctrl "github.com/arfan21/getprint-media/app/controllers/http"
	repo "github.com/arfan21/getprint-media/app/repository/mysql"
	"github.com/arfan21/getprint-media/app/services"
	"github.com/arfan21/getprint-media/configs/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(mySqlClient mysql.Client) *echo.Echo{
	route := echo.New()
	route.Use(middleware.Recover())
	route.Use(middleware.Logger())
	apiV1 := route.Group("/v1")

	// Routing Media
	mediaRepo := repo.NewMediaRepository(mySqlClient)
	mediaSrv := services.NewMediaServices(mediaRepo)
	mediaCtrl := ctrl.NewMediaControllers(mediaSrv)

	apiMedia := apiV1.Group("/media")
	apiMedia.POST("", mediaCtrl.Create)
	apiMedia.DELETE("", mediaCtrl.Delete)

	return route
}
