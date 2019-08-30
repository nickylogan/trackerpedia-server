package main

import (
	"log"
	"net/http"

	"trackerpedia-server/service"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", service.Home)
	router.HandleFunc("/tracker", service.UpdateStatusDelivery).Methods("POST")
	router.HandleFunc("/tracker/{id}", service.GetStatusDeliveryByID).Methods("GET")
	router.HandleFunc("/order", service.UpdateStatusOrder).Methods("POST")
	router.HandleFunc("/order/{id}", service.GetStatusOrderByID).Methods("GET")
	router.HandleFunc("/newDelivery", service.CreateNewDelivery).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
