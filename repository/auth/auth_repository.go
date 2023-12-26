package auth

import (
	modelAuth "app/model/auth"

	"gorm.io/gorm"
)

type AuthRepository interface {
	InsertUser(user *modelAuth.User) error
	GetUserByEmail(email string) (*modelAuth.User, error)
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{db}
}

func (ari *authRepositoryImpl) InsertUser(user *modelAuth.User) error {
	if err := ari.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ari *authRepositoryImpl) GetUserByEmail(email string) (*modelAuth.User, error) {
	var user modelAuth.User
	err := ari.db.Where("email = ?", email).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}
