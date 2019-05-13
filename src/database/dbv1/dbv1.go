package dbv1

import (
	"database/sql"
	"fmt"

	"google.golang.org/appengine"
)

// DBV1 represents the 1st version of the database
type DBV1 struct {
	dbc *sql.DB
}

// ConnectionDatabase connect the current database with the given credentials
func (db *DBV1) ConnectionDatabase(user, pass, dbname, url string) error {
	var err error

	if appengine.IsDevAppServer() {
		db.dbc, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp([127.0.0.1]:3306)/%s?parseTime=true", user, pass, dbname))
	} else {
		db.dbc, err = sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/%s?parseTime=true", user, pass, url, dbname))
	}

	if err == nil {
		err = db.dbc.Ping()
	}
	return err
}

// IsNewDB returns true when the database is no initialized yet
func (db *DBV1) IsNewDB() bool {
	return db.dbc == nil
}

// IsOldDB returns true when the database is already initialized
func (db *DBV1) IsOldDB() bool {
	return db.dbc != nil
}
