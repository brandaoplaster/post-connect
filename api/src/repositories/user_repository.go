package repositories

import (
	"api/api/src/models"
	"database/sql"
	"fmt"
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

func (repositiry Users) SearchByNameOrNickname(nameOrNickname string) ([]models.User, error) {
	nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname)

	result, erro := repositiry.database.Query(
		"select id, name, nickname, email, created_at from users where name LIKE ? or nickname LIKE ?",
		nameOrNickname, nameOrNickname,
	)

	if erro != nil {
		return nil, erro
	}
	defer result.Close()

	var users []models.User

	for result.Next() {
		var user models.User

		if erro = result.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}
