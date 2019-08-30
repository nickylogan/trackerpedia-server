package types

import "time"

// StatusDelivery is struct for updating delivery status every city
type StatusDelivery struct {
	IDResi int    `json:"IDResi"`
	Kota   string `json:"Kota"`
	Status int    `json:"Status"`
}

//StatusOrder is struct for updating order status
type StatusOrder struct {
	IDResi int    `json:"IDResi"`
	Status string `json:"Status"`
}

// Delivery is struct for Delivery tabel
type Delivery struct {
	IDResi  int       `json:"IDResi"`
	IDKota  int       `json:"IDKota"`
	Status  int       `json:"Status"`
	Ordinal int       `json:"Ordinal"`
	Time    time.Time `json:"Time"`
}

// Order is struct for Order Table
type Order struct {
	IDOrder int    `json:"IDOrder"`
	IDItem  int    `json:"IDItem"`
	Status  string `json:"Status"`
}
