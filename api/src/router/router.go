package router

import (
	"api/api/src/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	router := mux.NewRouter()
	return routes.ConfigureRoutes(router)
}
