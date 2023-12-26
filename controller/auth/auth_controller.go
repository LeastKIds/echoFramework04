package auth

import (
	repository "app/repository/auth"
	service "app/service/auth"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthController(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	authService := service.NewAuthServiceImpl(
		repository.NewAuthRepositoryImpl(db), validate)

	authRouter := e.Group("/auth")

	authRouter.POST("/sign-in", authService.SignIn)
	authRouter.POST("/sign-up", authService.SignUp)
	authRouter.POST("/logout", authService.Logout)

	authRouter.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "auth ok!!3")
	})
}
