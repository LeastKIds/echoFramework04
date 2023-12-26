package auth

import (
	model "app/model/auth"
	repository "app/repository/auth"
	"app/util/jwt"
	vo "app/vo/auth"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

	user := model.User{}
	copier.Copy(&user, &signUpRequest)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Server error")
	}
	user.Password = string(hashPassword)

	if err := a.ar.InsertUser(&user); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	signUpResponse := vo.SignUpResponse{}
	copier.Copy(&signUpResponse, &user)

	accessToken, err := jwt.TokenConfig(&user, c)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	signUpResponse.AccessToken = accessToken

	return c.JSON(http.StatusOK, signUpResponse)
}

func (a *authServiceImpl) Logout(c echo.Context) error {
	return nil
}
