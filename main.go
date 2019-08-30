package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type statusTracker struct {
	IDResi int32  `json:"IDResi"`
	Kota   string `json:"Kota"`
	Status int32  `json:"Status"`
}

type statusOrder struct {
	IDResi int    `json:"IDResi"`
	Status string `json:"Status"`
}

// Home is function on welcome page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Tracker System!")
}

// UpdateStatusTracker is function to update status tracker
func UpdateStatusTracker(w http.ResponseWriter, r *http.Request) {
	var response statusTracker

	trackerStatus, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(w, "[UpdateStatusTracker] failed to read the json body")
		return
	}
	err = json.Unmarshal(trackerStatus, &response)
	if err != nil {
		fmt.Fprintf(w, "[UpdateStatusTracker] failed when unmarshalling json")
		return
	}

	json.NewEncoder(w).Encode(response)
}

// UpdateStatusOrder is function to update status order
func UpdateStatusOrder(w http.ResponseWriter, r *http.Request) {
	var response statusOrder

	orderStatus, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(w, "[UpdateStatusOrder] failed to read the json body")
		return
	}

	err = json.Unmarshal(orderStatus, &response)
	if err != nil {
		fmt.Fprintf(w, "[UpdateStatusOrder] failed when unmarshalling json")
		return
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Home)
	router.HandleFunc("/tracker", UpdateStatusTracker).Methods("POST")
	router.HandleFunc("/order", UpdateStatusOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
