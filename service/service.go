package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"trackerpedia-server/config"
	"trackerpedia-server/types"

	"github.com/gorilla/mux"
)

type allDelivery []types.Delivery

var deliverys = allDelivery{
	{
		IDResi:  1,
		IDKota:  3,
		Status:  0,
		Ordinal: 1,
	},
	{
		IDResi:  1,
		IDKota:  1,
		Status:  1,
		Ordinal: 2,
	},
}

type allOrder []types.Order

var orders = allOrder{
	{
		IDOrder: 1,
		IDItem:  1,
		Status:  "Dalam Pengiriman",
	},
	{
		IDOrder: 2,
		IDItem:  3,
		Status:  "Di Gudang",
	},
}

// Home is function on welcome page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Tracker System!")
}

//GetStatusDeliveryByID is function to get all status tracker with parameter ID
func GetStatusDeliveryByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	idStatus, err := strconv.Atoi(param)
	if err != nil {
		fmt.Fprintf(w, "[GetStatusDeliveryByID] failed to read the json body")
		return
	}

	connect, err := config.Connection()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return
	}

	rows, err := db.Query("SELECT * from tb_delivery WHERE id_resi = $1", idStatus)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	var datas []types.Delivery

	for rows.Next() {
		data := types.Delivery{}
		err := rows.Scan(
			&data.IDResi,
			&data.IDKota,
			&data.Status,
			&data.Ordinal,
			&data.Time,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		datas = append(datas, data)
	}

	json.NewEncoder(w).Encode(datas)

}

// UpdateStatusDelivery is function to update status tracker
func UpdateStatusDelivery(w http.ResponseWriter, r *http.Request) {
	var response types.StatusDelivery

	deliveryStatus, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(w, "[UpdateStatusDelivery] failed to read the json body")
		return
	}
	err = json.Unmarshal(deliveryStatus, &response)
	if err != nil {
		fmt.Fprintf(w, "[UpdateStatusDelivery] failed when unmarshalling json")
		return
	}

	json.NewEncoder(w).Encode(response)
}

//GetStatusOrderByID is function to get all status tracker with parameter ID
func GetStatusOrderByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	idStatus, err := strconv.Atoi(param)
	if err != nil {
		fmt.Fprintf(w, "[GetStatusOrderByID] failed to read the json body")
		return
	}
	for _, oneOrder := range orders {
		if oneOrder.IDOrder == idStatus {
			json.NewEncoder(w).Encode(oneOrder)
		}
	}
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
