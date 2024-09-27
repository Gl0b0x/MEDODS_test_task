package repository

import (
	"MEDODS/pkg/models"
	"database/sql"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (repo *AuthPostgres) GetUserByUUID(uuid string) (*models.User, error) {
	var user models.User
	q := "SELECT * FROM users WHERE uuid=$1"
	row := repo.db.QueryRow(q, uuid)
	err := row.Scan(&user.Guid, &user.Email, &user.Ip, &user.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *AuthPostgres) UpdateUserRefreshToken(uuid, hashRefreshToken string) error {
	q := "UPDATE users SET refresh_token = $1 WHERE uuid = $2"
	_, err := repo.db.Exec(q, hashRefreshToken, uuid)
	return err
}
