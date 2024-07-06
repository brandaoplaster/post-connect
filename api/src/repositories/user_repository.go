package repositories

import (
	"api/api/src/models"
	"database/sql"
)

type Users struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) *Users {
	return &Users{database}
}

func (repositiry Users) Create(user models.User) (uint64, error) {
	statement, erro := repositiry.database.Prepare(
		"insert into users (name, nickname, email, password) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.NickName, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastInsert, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsert), nil

}
