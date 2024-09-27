package repository

import (
	"MEDODS/pkg/models"
	"database/sql"
)

type Authorization interface {
	GetUserByUUID(uuid string) (*models.User, error)
	UpdateUserRefreshToken(uuid, hashRefreshToken string) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
