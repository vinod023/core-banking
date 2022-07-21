package config

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ConfigRouter() *echo.Echo {
	router := echo.New()
	router.HideBanner = false
	router.HidePort = false
	router.Debug = true

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	//CORS
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	return router
}
