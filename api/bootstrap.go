package main

import (
    "api/api/src/controllers"
    "api/api/src/repositories"
    "api/api/src/router/routes"
    "database/sql"
)

func NewHandlers(db *sql.DB) *routes.Handlers {
    userRepo := repositories.NewUserRepository(db)

    return &routes.Handlers{
        Users: controllers.NewUsersController(userRepo),
    }
}
