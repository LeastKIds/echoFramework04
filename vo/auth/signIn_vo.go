package auth

import "time"

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=16"`
}

type SignInResponse struct {
	ID          uint      `json:"id"`
	Email       string    `json:"email"`
	CreatedUp   time.Time `json:"created_up"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessToken string    `json:"access_token"`
}
