package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CRUD --> CREATE, REAR, UPDATE, DELETE

// CREATE --> POST
// READ --> GET
// UPDATE --> PUT
// DELETE --> DELETE

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

	fmt.Println("Escurando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
