package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"trackerpedia-server/types"
)

// Home is function on welcome page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Tracker System!")
}

// UpdateStatusTracker is function to update status tracker
func UpdateStatusTracker(w http.ResponseWriter, r *http.Request) {
	var response types.StatusTracker

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
	var response types.StatusOrder

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
