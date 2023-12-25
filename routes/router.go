package routes

import (
	"app/config"
	controllerAuth "app/controller/auth"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func NewRouter() (*echo.Echo, config.EnvConfig) {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	var cfg config.EnvConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	controllerAuth.AuthController(e)

	return e, cfg
}
