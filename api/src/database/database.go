package database

import (
	"api/api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	database, erro := sql.Open("mysql", config.APIURL)

	if erro = database.Ping(); erro != nil {
		database.Close()
		return nil, erro
	}

	return database, nil
}
