package postgres

import (
	"database/sql"
	"fmt"
)

func New(dbName, host, port, userDB, password, driverName, sslmode string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, userDB, password, dbName, sslmode)
	db, err := sql.Open(driverName, psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
