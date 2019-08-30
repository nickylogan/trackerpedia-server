package types

import "time"

// Delivery is struct for Delivery tabel
type Delivery struct {
	IDResi  int       `json:"idResi"`
	Kota    string    `json:"city"`
	Status  int       `json:"status"`
	Ordinal int       `json:"ordinal"`
	Time    time.Time `json:"time"`
}

// Order is struct for Order Table
type Order struct {
	IDOrder            int       `json:"idOrder"`
	NamaItem           string    `json:"nameItem"`
	Weight             string    `json:"weight"`
	Status             string    `json:"status"`
	Time               time.Time `json:"time"`
	DestinationAddress string    `json:"destinationAddress"`
	DestinationCity    string    `json:"destinationCity"`
}
