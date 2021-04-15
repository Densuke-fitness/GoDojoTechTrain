package main

import (
	"log"
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	const port string = ":8080"

	router.HandleFunc("/users/create", handler.CreateUser()).Methods("POST")
	log.Println("Server listeining on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
