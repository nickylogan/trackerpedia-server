package types

import "time"

// Delivery is struct for Delivery tabel
type Delivery struct {
	IDResi  int       `json:"idResi"`
	Kota    string    `json:"city"`
	Status  int       `json:"status"`
	Ordinal int       `json:"ordinal"`
	Time    time.Time `json:"time"`
	Next    bool      `json:"next"`
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

type Item struct {
	IDItem   int    `json:"idItem`
	NameItem string `json:"nameItem"`
	Price    int    `json:"price"`
	Weight   int    `json:"weight"`
}
