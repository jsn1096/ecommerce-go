package main

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newHTTP(errorHandler echo.HTTPErrorHandler) *echo.Echo {
	e := echo.New()

	// Loguea todas las peticiones
	e.Use(middleware.Logger())
	// Recuperarse de un error que pueda romper el sistema
	e.Use(middleware.Recover())

	// Permitir los cors
	corsConfig := middleware.CORSConfig{
		AllowOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		AllowMethods: strings.Split(os.Getenv("ALLOWED_METHODS"), ","),
	}

	e.Use(middleware.CORSWithConfig(corsConfig))

	e.HTTPErrorHandler = errorHandler
	return e
}
