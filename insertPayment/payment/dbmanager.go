package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBCONN struct {
	DBconn *sql.DB
}

func DBConnection() (DBCONN, error) {
	db, err := sql.Open("mysql", "vegas:!mgrsol123@tcp(int.trustnhope.com:6306)/h01042")
	if err != nil {
		log.Fatal(err)
		return DBCONN{}, err
	}

	return DBCONN{
		DBconn: db,
	}, nil

}
