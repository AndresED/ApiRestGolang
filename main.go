package main
import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
)
func main(){
	router:=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",Index)
	router.HandleFunc("/contacto",Contact)
	router.HandleFunc("/movie",MovieList)
	router.HandleFunc("/movie/{id}",MovieShow)
	server:= http.ListenAndServe(":8080",router)
	log.Fatal(server)
	fmt.Println("Servidor ejecutandose en http://localhost:8080")
}

func Index(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hola mundo desde mi servidor web con Go")
}

func Contact(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Pagina de contacto")
}


func MovieShow(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	movie_id:=params["id"]
	fmt.Fprintf(w,"Detalle de la pelicula %s",movie_id)	
}

func MovieList(w http.ResponseWriter,r *http.Request){
	movies:=Movies{
		{"Pelicula 1","2017","director 1"},
		{"Pelicula 2","2017","director 2"},
		{"Pelicula 3","2017","director 3"}}
	json.NewEncoder(w).Enconde(movies)
}