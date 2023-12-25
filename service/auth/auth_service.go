package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	SignUp(c echo.Context) error
	SignIn(c echo.Context) error
	Logout(c echo.Context) error
}

type authServiceImpl struct {
}

func NewAuthServiceImpl() AuthService {
	return &authServiceImpl{}
}

func (a *authServiceImpl) SignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, "test ok!")
}

func (a *authServiceImpl) SignUp(c echo.Context) error {
	return nil
}

func (a *authServiceImpl) Logout(c echo.Context) error {
	return nil
}
