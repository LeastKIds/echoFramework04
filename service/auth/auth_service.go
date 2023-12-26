package auth

import (
	repository "app/repository/auth"
	vo "app/vo/auth"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AuthService interface {
	SignUp(c echo.Context) error
	SignIn(c echo.Context) error
	Logout(c echo.Context) error
}

type authServiceImpl struct {
	ar       repository.AuthRepository
	validate *validator.Validate
}

func NewAuthServiceImpl(ar repository.AuthRepository, validate *validator.Validate) AuthService {
	return &authServiceImpl{ar, validate}
}

func (a *authServiceImpl) SignIn(c echo.Context) error {

	return c.JSON(http.StatusOK, "test ok!dafafdasfd1231312312321a")
}

func (a *authServiceImpl) SignUp(c echo.Context) error {
	signUpRequest := vo.SignUpRequest{}
	if err := c.Bind(&signUpRequest); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest, "json format error")
	}

	if err := a.validate.Struct(signUpRequest); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest, "json format error")
	}

	return c.JSON(http.StatusOK, signUpRequest)
}

func (a *authServiceImpl) Logout(c echo.Context) error {
	return nil
}
