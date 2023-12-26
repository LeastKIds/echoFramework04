package jwt

import (
	model "app/model/auth"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func TokenConfig(user *model.User, c echo.Context) (string, error) {
	accessToken, err := CreateAccessToken(user)
	if err != nil {
		c.Logger().Error(err.Error())
		return "", err
	}

	refreshToken, err := CreateRefreshToken(user)
	if err != nil {
		c.Logger().Error(err.Error())
		return "", err
	}

	// accessTokenCookie := new(http.Cookie)
	// accessTokenCookie.Name = "session_token"
	// accessTokenCookie.Value = accessToken
	// accessTokenCookie.Expires = time.Now().Add(30 * time.Minute)
	// c.SetCookie(accessTokenCookie)

	refreshTokenCookie := new(http.Cookie)
	refreshTokenCookie.Name = "refresh_token"
	refreshTokenCookie.Value = refreshToken
	refreshTokenCookie.HttpOnly = true
	refreshTokenCookie.Secure = true
	refreshTokenCookie.Path = "/"
	refreshTokenCookie.Expires = time.Now().Add(60 * 24 * 14 * time.Minute)
	c.SetCookie(refreshTokenCookie)

	return accessToken, nil

}

func CreateAccessToken(user *model.User) (string, error) {
	return createToken(user, 30)
}

func CreateRefreshToken(user *model.User) (string, error) {
	return createToken(user, 60*24*14)
}

func createToken(user *model.User, expTime int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Duration(expTime) * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
