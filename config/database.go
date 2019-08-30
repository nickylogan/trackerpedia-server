package config

import (
	"database/sql"
	"fmt"
)

const (
	host     = "10.50.218.7"
	port     = 5432
	user     = "devcamp"
	password = "devcamp2019"
	dbname   = "devcamp"
)

// Connection is function for connecting system to database
func Connection() (string, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return psqlInfo, nil
}
