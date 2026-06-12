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


func (repositiry Users) Create(user models.User) (models.User, error) {
	statement, erro := repositiry.database.Prepare(
		"insert into users (name, nickname, email, password) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.NickName, user.Email, user.Password)
	if erro != nil {
		return models.User{}, erro
	}

	lastInsert, erro := result.LastInsertId()
	if erro != nil {
		return models.User{}, erro
	}

	user.ID = uint64(lastInsert)
	user.Password = ""

	return user, nil
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

func (repositiry Users) FindById(id uint64) (models.User, error) {
	result, erro := repositiry.database.Query(
		"select id, name, nickname, email, created_at, updated_at from users where id = ?",
		id,
	)

	if erro != nil {
		return models.User{}, erro
	}

	defer result.Close()

	var user models.User

	if result.Next() {
		if erro = result.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}
	return user, nil
}

func (repository Users) All() ([]models.User, error) {
	result, erro := repository.database.Query(
		"select id, name, nickname, email, created_at, updated_at from users",
	)

	if erro != nil {
		return []models.User{}, erro
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
			&user.UpdatedAt,
		); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository Users) Update(id uint64, user models.User) (models.User, error) {
	statement, erro := repository.database.Prepare(
		"update users set name = ?, nickname = ?, email = ? where id = ?",
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.NickName, user.Email, id)
	if erro != nil {
		return models.User{}, erro
	}

	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return models.User{}, erro
	}

	if rowsAffected == 0 {
		return models.User{}, fmt.Errorf("user not found")
	}

	return repository.FindById(id)
}

func (repository Users) Delete(id uint64) error {
	statement, erro := repository.database.Prepare(
		"delete from users where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	result, erro := statement.Exec(id)
	if erro != nil {
		return erro
	}

	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return  erro
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
