package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JairDavid/controller"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	r.HandleFunc("/", controller.Index).Methods("GET")
	r.HandleFunc("/technologies", controller.Technologies).Methods("GET")
	r.HandleFunc("/probien", controller.Probien).Methods("GET")
	r.HandleFunc("/gespro", controller.Gespro).Methods("GET")
	r.HandleFunc("/covec", controller.Covec).Methods("GET")
	r.HandleFunc("/pibe", controller.Pibe).Methods("GET")

	server := &http.Server{Addr: ":" + port, Handler: r, IdleTimeout: time.Second * 5}

	log.Fatal(server.ListenAndServe())
}
