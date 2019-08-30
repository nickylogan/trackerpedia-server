package main

import (
	"log"
	"net/http"

	"trackerpedia-server/service"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", service.Home)
	router.HandleFunc("/tracker", service.UpdateStatusTracker).Methods("POST")
	router.HandleFunc("/order", service.UpdateStatusOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
