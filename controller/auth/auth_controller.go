package auth

import (
	"app/service/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthController(e *echo.Echo) {
	authService := auth.NewAuthServiceImpl()
	authRouter := e.Group("/auth")

	authRouter.POST("/sign-in", authService.SignIn)
	authRouter.POST("/sign-up", authService.SignUp)
	authRouter.POST("/logout", authService.Logout)

	authRouter.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "auth ok!!")
	})
}
