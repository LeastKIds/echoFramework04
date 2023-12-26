package auth

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedUp time.Time `json:"created_up"`
	UpdatedAt time.Time `json:"updated_at"`
}
