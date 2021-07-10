package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRoute() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes{
	Route{"index", "GET", "/", Index},
	Route{"contacto", "GET", "/contacto", Contacto},
	Route{"peliculasList", "GET", "/peliculas", peliculasList},
	Route{"peliculasShow", "GET", "/peliculas/{id}", peliculasShow},
}
