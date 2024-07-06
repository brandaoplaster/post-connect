package repositories

import "database/sql"

type Users struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) *Users {
	return &Users{database}
}
