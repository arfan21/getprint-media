package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arfan21/getprint-media/configs"
	_mediaCtrl "github.com/arfan21/getprint-media/controllers/http/media"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	route := echo.New()
	route.Use(middleware.Recover())
	route.Use(middleware.Logger())

	db, err := configs.MysqlConnect()
	if err != nil {
		route.Logger.Fatal(err)
	}

	mediaCtrl := _mediaCtrl.NewMediaControllers(db)
	mediaCtrl.Routes(route)

	filename := "tes.exe"
	extensionFile := filepath.Ext(filename)
	fmt.Println(extensionFile)

	route.Logger.Fatal(route.Start(fmt.Sprintf(":%s", port)))
}
