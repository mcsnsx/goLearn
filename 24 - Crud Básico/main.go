package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CRUD --> CREATE, REAR, UPDATE, DELETE

func main() {
	router := mux.NewRouter()

	fmt.Println("Escurando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
