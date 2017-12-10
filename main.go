package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", Contact)
	router.HandleFunc("/movie", MovieList)
	router.HandleFunc("/movie/{id}", MovieShow)
	fmt.Println("Servidor ejecutandose en http://localhost:8080")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

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
		Movie{"Pelicula 1", 2017, "director 1"},
		Movie{"Pelicula 2", 2017, "director 2"},
		Movie{"Pelicula 3", 2017, "director 3"},
	}
	json.NewEncoder(w).Encode(movies)
}
