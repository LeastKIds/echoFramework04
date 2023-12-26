package db

import (
	"app/config"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.EnvConfig, e *echo.Echo) *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(postgres.Open(cfg.DataSourceName()), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	e.Logger.Info("database connected")
	return db

}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
