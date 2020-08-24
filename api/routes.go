package api

import (
	"./data"
	"./handlers"
	"github.com/gorilla/mux"
)

var routes = data.Routes{
	data.Route{
		Name:       "Index",
		Method:     "GET",
		Pattern:    "/auth",
		HandleFunc: handlers.IndexHandler,
	},
	data.Route{
		Name:       "Login",
		Method:     "POST",
		Pattern:    "/auth/login",
		HandleFunc: handlers.LoginHandler,
	},
	data.Route{
		Name:       "Login",
		Method:     "POST",
		Pattern:    "/auth/create",
		HandleFunc: handlers.UserCreateHandler,
	},
}

func NewRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		r.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}

	return r
}
