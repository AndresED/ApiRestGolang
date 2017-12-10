package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

var movies = Movies{
	Movie{name: "Pelicula 1", year: 2017, director: "director 1"},
	Movie{name: "Pelicula 2", year: 2017, director: "director 2"},
	Movie{name: "Pelicula 3", year: 2017, director: "director 3"},
}
var collection = getSesion().DB("pelicula_go").C("movies")

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
	fmt.Println(movies)
	json.NewEncoder(w).Encode(movies)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movieData Movie
	err := decoder.Decode(&movieData)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	log.Println(movieData)

	//realizando conexion a la bd
	err2 := collection.Insert(movieData)
	if err2 != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func getSesion() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session

}
