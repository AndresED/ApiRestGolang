package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := newRouter()
	fmt.Println("Servidor ejecutandose en http://localhost:8080")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
