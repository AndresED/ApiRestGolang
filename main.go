package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//CREANDO UN OBJETO QUE SE ENCARGARA DE TRABAJAR NUESTRAS RUTAS
	router := newRouter()
	//MENSAJE DE LENTAMIENTO DE SERVIDOR
	fmt.Println("Servidor ejecutandose en http://localhost:8080")
	//CREANDO SERVIDOR
	server := http.ListenAndServe(":8080", router)
	//SI HAY ALGUN ERROR NOS LO DIRA
	log.Fatal(server)
}
