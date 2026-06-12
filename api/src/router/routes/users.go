package routes

import (
	"net/http"
)

func userRoutes(handler *Handlers) []Route{
	return []Route{
		{
			URI:           "/users",
			Method:        http.MethodPost,
			Function:      handler.Users.Create,
			Authenticated: false,
		},
		{
			URI:           "/users",
			Method:        http.MethodGet,
			Function:      handler.Users.Index,
			Authenticated: false,
		},
		{
			URI:           "/users/{userId}",
			Method:        http.MethodGet,
			Function:      handler.Users.Show,
			Authenticated: false,
		},
		{
			URI:           "/users/{userId}",
			Method:        http.MethodPut,
			Function:      handler.Users.Update,
			Authenticated: false,
		},
		{
			URI:           "/users/{userId}",
			Method:        http.MethodDelete,
			Function:      handler.Users.Delete,
			Authenticated: false,
		},
	}
}
