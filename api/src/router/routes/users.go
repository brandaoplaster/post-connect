package routes

import (
	"api/api/src/router/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:           "/users",
		Method:        http.MethodPost,
		Function:      controllers.Create,
		Authenticated: false,
	},
	{
		URI:           "/users",
		Method:        http.MethodGet,
		Function:      controllers.Index,
		Authenticated: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodGet,
		Function:      controllers.Show,
		Authenticated: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodPut,
		Function:      controllers.Update,
		Authenticated: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodDelete,
		Function:      controllers.Delete,
		Authenticated: false,
	},
}
