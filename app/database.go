package main

import (
	"database/sql"
	"fmt"
	"log"

	"google.golang.org/appengine"
)

var db *sql.DB

func initDB(user, password, dbname, connectionName string) {
	var err error

	if appengine.IsDevAppServer() {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp([127.0.0.1]:3306)/%s", user, password, dbname))
	} else {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/%s", user, password, connectionName, dbname))
	}

	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}
}
