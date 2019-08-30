package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"trackerpedia-server/config"
	"trackerpedia-server/types"

	"github.com/gorilla/mux"
)

// Home is function on welcome page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Tracker System!")
}

//GetStatusDeliveryByID is function to get all status tracker with parameter ID
func GetStatusDeliveryByID(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
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

	rows, err := db.Query("SELECT id_resi, nama_city, status, ordinal, date_time from tb_delivery INNER JOIN tb_city ON tb_delivery.id_kota = tb_city.id_city WHERE id_resi = $1 AND status = 1", idStatus)
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
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	// StatusDelivery is struct for updating delivery status every city
	type statusDelivery struct {
		IDResi int `json:"idResi"`
		IDKota int `json:"idKota"`
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
	(w).Header().Set("Access-Control-Allow-Origin", "*")
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

	rows, err := db.Query("SELECT id_order,nama_item,weight,status,time_stamp,destination_address,destination_city from tb_order INNER JOIN tb_item ON tb_order.id_item=tb_item.id_item WHERE id_order = $1", idStatus)
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
			&data.NamaItem,
			&data.Weight,
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
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	//StatusOrder is struct for updating order status
	type statusOrder struct {
		IDOrder int    `json:"idOrder"`
		Status  string `json:"status"`
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

// CreateNewDelivery is function for creating new delivery
func CreateNewDelivery(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	type dataDelivery struct {
		IDResi int `json:"idResi"`
	}

	var response dataDelivery

	deliveryData, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(w, "[UpdateStatusOrder] failed to read the json body")
		return
	}

	err = json.Unmarshal(deliveryData, &response)
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

	var statusDelivery = 1

	for i := 1; i <= 5; i++ {
		if i > 1 {
			statusDelivery = 0
		}

		_, err = db.Exec("INSERT INTO tb_delivery (id_resi, id_kota, status, ordinal) VALUES ($1, $2, $3, $4)", response.IDResi, i, statusDelivery, i)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	_, err = db.Exec("UPDATE tb_order SET status = $1 WHERE id_order = $2", "SENT", response.IDResi)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//GetOrderSent is function for get All drom order where status is SENT
func GetOrderSent(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	connect, err := config.Connection()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return
	}

	rows, err := db.Query("SELECT id_order,nama_item,weight,status,time_stamp,destination_address,destination_city from tb_order INNER JOIN tb_item ON tb_order.id_item=tb_item.id_item WHERE tb_order.status = $1", "SENT")
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
			&data.NamaItem,
			&data.Weight,
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

// GetAllOrder is function for get All from table order
func GetAllOrder(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	connect, err := config.Connection()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return
	}

	rows, err := db.Query("SELECT id_order,nama_item,weight,status,time_stamp,destination_address,destination_city from tb_order INNER JOIN tb_item ON tb_order.id_item=tb_item.id_item")
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
			&data.NamaItem,
			&data.Weight,
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

// CreateNewOrder is function to create new order
func CreateNewOrder(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	type newOrder struct {
		IDItem  int    `json:"idItem"`
		Address string `json:"address"`
	}
	var response newOrder

	orderNew, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(w, "[UpdateStatusOrder] failed to read the json body")
		return
	}

	err = json.Unmarshal(orderNew, &response)
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

	idOrder := rand.Intn(10000)

	_, err = db.Exec("INSERT INTO tb_order (id_order, id_item, status, destination_address, destination_city) VALUES ($1, $2, $3, $4, $5)", idOrder, response.IDItem, "PENDING", response.Address, "Jakarta")
	if err != nil {
		fmt.Println(err)
		return
	}
}
