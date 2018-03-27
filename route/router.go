package route

import (
	"net/http"
	"rest-api/controller"

	"github.com/gorilla/mux"
)

// Route estructuras para api
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes .
type Routes []Route

// NewMuxRouter manejador de rutas
func NewMuxRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"GetUsers",
		"GET",
		"/users",
		controller.GetUsers,
	},
	Route{
		"GetUsers",
		"POST",
		"/users",
		controller.GetUsers,
	},
	Route{
		"GetUser",
		"GET",
		"/user/{id}",
		controller.GetUser,
	},
	/* 	Route{
		"UserDelete",
		"DELETE",
		"/user/{id}",
		controller.UserDelete,
	}, */
}
