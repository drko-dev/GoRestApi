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
		"/upload/{id}",
		controller.GetUser,
	},
	Route{
		"Upload",
		"POST",
		"/upload",
		controller.UploadFile,
	},
	Route{
		"AnalyzeFile",
		"GET",
		"/analyze/{image}",
		controller.AnalyzeFile,
	},
}

// NewMuxRouter manejador de rutas
func NewMuxRouter() *mux.Router {

	router := mux.NewRouter()
	ServeStatic(router, "./public/")
	RoutesServer(router)
	// log.Printf("%#v\n", router)
	return router
}

// RoutesServer servidor de rutas
func RoutesServer(router *mux.Router) {
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
}

// ServeStatic carpetas estaticas
func ServeStatic(router *mux.Router, staticDirectory string) {
	staticPaths := map[string]string{
		"templates": staticDirectory + "templates/",
		"css":       staticDirectory + "css/",
		"images":    staticDirectory + "images/",
		// "js":        staticDirectory + "js/",
	}
	for pathName, pathValue := range staticPaths {
		pathPrefix := "/" + pathName + "/"
		router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
			http.FileServer(http.Dir(pathValue))))
	}
}
