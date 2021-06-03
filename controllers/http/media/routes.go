package media

import "github.com/labstack/echo/v4"

func (ctrl *mediaControllers) Routes(route *echo.Echo) {
	route.POST("/media", ctrl.Create)
	route.DELETE("/media/:id", ctrl.Delete)
}
