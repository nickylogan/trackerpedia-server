package types

// StatusTracker is struct for updating tracker status every city
type StatusTracker struct {
	IDResi int32  `json:"IDResi"`
	Kota   string `json:"Kota"`
	Status int32  `json:"Status"`
}

//StatusOrder is struct for
type StatusOrder struct {
	IDResi int    `json:"IDResi"`
	Status string `json:"Status"`
}
