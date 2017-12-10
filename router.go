package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc).
			Name(route.Name)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Contacto",
		"GET",
		"/contacto",
		Contact,
	},
	Route{
		"ListaPeliculas",
		"GET",
		"/movie",
		MovieList,
	},
	Route{
		"DetallePelicula",
		"GET",
		"/movie/{id}",
		MovieShow,
	},
	Route{
		"MovieADD",
		"POST",
		"/add-pelicula",
		MovieAdd,
	},
}
