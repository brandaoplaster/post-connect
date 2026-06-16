package main

import (
    "api/api/src/controllers"
    "api/api/src/repositories"
    "api/api/src/router/routes"
    "api/api/src/services"
    "database/sql"
)

func NewHandlers(db *sql.DB) *routes.Handlers {
    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)

    return &routes.Handlers{
        Users: controllers.NewUsersController(userService),
    }
}
