package main

import (
	"log"
	"net/http"

	"github.com/Densuke-fitness/GoDojoTechTrain/controller"
	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	const port string = ":8080"
	defer dbConnection.GetInstance().Close()

	router.HandleFunc("/user/create", controller.CreateUser()).Methods("POST")
	router.HandleFunc("/user/get", controller.GetUser()).Methods("GET")
	router.HandleFunc("/user/update", controller.UpdateUser()).Methods("PUT")
	router.HandleFunc("/gacha/draw", controller.DrawGacha()).Methods("POST")
	router.HandleFunc("/character/list", controller.GetCharacterList()).Methods("GET")
	log.Println("Server listeining on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
