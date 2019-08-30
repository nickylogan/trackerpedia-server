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
		Kota:    "Jepara",
		Status:  0,
		Ordinal: 1,
	},
	{
		IDResi:  1,
		Kota:    "Jakarta",
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

	rows, err := db.Query("SELECT id_resi, nama_city, status, ordinal, date_time from tb_delivery INNER JOIN tb_city ON tb_delivery.id_kota = tb_city.id_city WHERE id_resi = $1", idStatus)
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
			&data.Kota,
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

	// StatusDelivery is struct for updating delivery status every city
	type statusDelivery struct {
		IDResi int `json:"IDResi"`
		IDKota int `json:"IDKota"`
	}
	var response statusDelivery

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

	connect, err := config.Connection()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("postgres", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("UPDATE tb_delivery SET status = 1 WHERE id_resi = $1 AND id_kota = $2", response.IDResi, response.IDKota)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//GetStatusOrderByID is function to get all status tracker with parameter ID
func GetStatusOrderByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	idStatus, err := strconv.Atoi(param)
	if err != nil {
		fmt.Fprintf(w, "[GetStatusOrderByID] failed to read the json body")
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

	rows, err := db.Query("SELECT id_order,id_item,status,time_stamp,destination_address,destination_city from tb_order WHERE id_order = $1", idStatus)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	var datas []types.Order

	for rows.Next() {
		data := types.Order{}
		err := rows.Scan(
			&data.IDOrder,
			&data.IDItem,
			&data.Status,
			&data.Time,
			&data.DestinationAddress,
			&data.DestinationCity,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		datas = append(datas, data)
	}

	json.NewEncoder(w).Encode(datas)
}

// UpdateStatusOrder is function to update status order
func UpdateStatusOrder(w http.ResponseWriter, r *http.Request) {

	//StatusOrder is struct for updating order status
	type statusOrder struct {
		IDOrder int    `json:"IDOrder"`
		Status  string `json:"Status"`
	}
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

	connect, err := config.Connection()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("postgres", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("UPDATE tb_order SET status = $1 WHERE id_order = $2", response.Status, response.IDOrder)
	if err != nil {
		fmt.Println(err)
		return
	}
}
