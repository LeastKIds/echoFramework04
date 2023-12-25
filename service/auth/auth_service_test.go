package auth

import (
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	c := http.Client{}
	r, err := c.Post("http://localhost:8080/auth/sign-in", "application/json", strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, r.StatusCode, "X")
}
