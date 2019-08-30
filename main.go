package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"trackerpedia-server/service"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", service.Home)
	router.HandleFunc("/tracker", service.UpdateStatusDelivery).Methods("POST", "OPTIONS")
	router.HandleFunc("/tracker/{id}", service.GetStatusDeliveryByID).Methods("GET")
	router.HandleFunc("/order", service.UpdateStatusOrder).Methods("POST", "OPTIONS")
	router.HandleFunc("/order/{id}", service.GetStatusOrderByID).Methods("GET")
	router.HandleFunc("/order_sent", service.GetOrderSent).Methods("GET")
	router.HandleFunc("/allOrder", service.GetAllOrder).Methods("GET")
	router.HandleFunc("/newDelivery", service.CreateNewDelivery).Methods("POST", "OPTIONS")
	router.HandleFunc("/newOrder", service.CreateNewOrder).Methods("POST", "OPTIONS")
	router.HandleFunc("/allItem", service.GetAllItem).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
