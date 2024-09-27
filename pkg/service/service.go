package service

import (
	"MEDODS/pkg/models"
	"MEDODS/pkg/repository"
)

type Authorization interface {
	GetUserByUUID(uuid string) (*models.User, error)
	UpdateUserRefreshToken(user *models.User, hashRefreshToken string) error
	generateAccessToken(uuid, ip string) (string, error)
	generateRefreshToken(uuid, ip string) (string, error)
	ParseToken(tokenString string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
	GenerateTokens(user *models.User) (string, string, error)
}
type Service struct {
	Authorization
}

func NewService(repos *repository.Repository, jwtSecret string) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, jwtSecret),
	}
}
