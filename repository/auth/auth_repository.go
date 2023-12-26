package auth

import (
	modelAuth "app/model/auth"

	"gorm.io/gorm"
)

type AuthRepository interface {
	InsertUser(user *modelAuth.User) (*modelAuth.User, error)
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{db}
}

func (ari *authRepositoryImpl) InsertUser(user *modelAuth.User) (*modelAuth.User, error) {
	if err := ari.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
