package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con Go")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pagina de contacto")
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Detalle de la pelicula %s", movie_id)
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{name: "Pelicula 1", year: 2017, director: "director 1"},
		Movie{name: "Pelicula 2", year: 2017, director: "director 2"},
		Movie{name: "Pelicula 3", year: 2017, director: "director 3"},
	}
	fmt.Println(movies)
	json.NewEncoder(w).Encode(movies)
}
