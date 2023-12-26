package auth

import (
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	c := http.Client{}
	r, err := c.Post("http://localhost:8080/auth/sign-up", "application/json", strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}
	responseBody, _ := io.ReadAll(r.Body)
	assert.Equal(t, http.StatusBadRequest, r.StatusCode, string(responseBody))

	r2, err := c.Post("http://localhost:8080/auth/sign-up", "application/json", strings.NewReader(`{"Email": "test", "Password": "123"}`))
	if err != nil {
		log.Fatal(err)
	}
	responseBody2, _ := io.ReadAll(r2.Body)
	assert.Equal(t, http.StatusBadRequest, r2.StatusCode, string(responseBody2))

	r3, err := c.Post("http://localhost:8080/auth/sign-up", "application/json", strings.NewReader(`{"Email": "test@naver.com", "Password": "123"}`))
	if err != nil {
		log.Fatal(err)
	}
	responseBody3, _ := io.ReadAll(r3.Body)
	assert.Equal(t, http.StatusBadRequest, r3.StatusCode, string(responseBody3))

	r4, err := c.Post("http://localhost:8080/auth/sign-up", "application/json", strings.NewReader(`{"Email": "test@naver.com", "Password": "12345"}`))
	if err != nil {
		log.Fatal(err)
	}
	responseBody4, _ := io.ReadAll(r4.Body)
	assert.Equal(t, http.StatusOK, r4.StatusCode, string(responseBody4))

}

func TestSignUp(t *testing.T) {

}
